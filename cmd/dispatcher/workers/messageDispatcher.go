package workers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	ppt "github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/deyvidm/sms/cmd/dispatcher/client"
	"github.com/deyvidm/sms/cmd/dispatcher/log"
	"github.com/deyvidm/sms/pkg/tasks"
	"github.com/deyvidm/sms/pkg/types"
	"github.com/deyvidm/sms/pkg/utils"
	"github.com/hibiken/asynq"
)

var logger = log.GetLogger()

type MessageDispatcher struct {
	wbc *client.WebBackendClient
	irs *client.InviteResponseStore
	pp  *pinpoint.Client
}

func NewMessageDispatcher(wbc *client.WebBackendClient, irs *client.InviteResponseStore, pp *pinpoint.Client) *MessageDispatcher {
	return &MessageDispatcher{
		wbc: wbc,
		irs: irs,
		pp:  pp,
	}
}

func (md *MessageDispatcher) HandleSendInviteTask(ctx context.Context, t *asynq.Task) error {
	var p tasks.NewInvitePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	logger.Infof("|%s|\tinviting %s to %s : '%s'", t.Type(), p.ToPhoneNumber, p.InviteID, p.Content)

	if err := md.irs.SaveNewInviteEntry(p.ToPhoneNumber, p.InviteID); err != nil {
		return err
	}

	return md.wbc.UpdateInvite(client.UpdateInvite{
		ID:     p.InviteID,
		Status: utils.Ptr(types.InviteStatus_Invited.String()),
	})
}

func (md *MessageDispatcher) HandleNewMessageTask(ctx context.Context, t *asynq.Task) error {
	var p tasks.NewMessagePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	logger.Infof("|%s|\tsending one-off message to %s : '%s'", t.Type(), p.ToPhoneNumber, p.Content)

	resp, err := md.pp.SendMessages(ctx, &pinpoint.SendMessagesInput{
		ApplicationId: utils.Ptr("ecea11cc234a4af78bfe9831beca48bf"),
		MessageRequest: &ppt.MessageRequest{
			Addresses: map[string]ppt.AddressConfiguration{
				p.ToPhoneNumber: {ChannelType: ppt.ChannelTypeSms},
			},
			MessageConfiguration: &ppt.DirectMessageConfiguration{
				SMSMessage: &ppt.SMSMessage{
					Body:        &p.Content,
					MessageType: ppt.MessageTypePromotional,
				},
			},
		},
	})

	if err != nil {
		logger.Infof("|%s|\t fucked up sending, returning error", t.Type())
		return err
	}

	logger.Infof("|%s|\t finished sending message", t.Type())
	logger.Infof("|%s|\t dumping Result Metadata:", t.Type())
	logger.Info(utils.JSONDump(resp.ResultMetadata))

	logger.Infof("|%s|\t dumping Message Response:", t.Type())
	logger.Info(utils.JSONDump(resp.MessageResponse))
	return nil
}
