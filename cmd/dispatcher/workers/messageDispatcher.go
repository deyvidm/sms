package workers

import (
	"context"
	"encoding/json"
	"fmt"

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

	// pushes an invite onto an attendee's stack
	if err := md.irs.SaveNewInviteEntry(p.ToPhoneNumber, p.InviteID); err != nil {
		return err
	}

	// TOOD actually send an SMS

	// tell the web-server that we've sent an invite (ie. update status from Pending to Sent or whatever)
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

	if len(p.Content) < 1 {
		err := fmt.Errorf("message request contains no body")
		logger.Error(err)
		return err
	}

	// 10-digit phone number + area code (area code will always be "+X" where X is at least a single digit number)
	// this makes the total number of characters in a valid number at least 12 (plus sign included)
	// ex: +1xxxiiizzzz for North America => 12 characters
	// ex: +41xxxiiizzzz for Switzerland  => 13 characters
	if len(p.ToPhoneNumber) < 12 {
		err := fmt.Errorf("unexpected phone number length: %d [%s]", len(p.ToPhoneNumber), p.ToPhoneNumber)
		logger.Error(err)
		return err
	}

	logger.Infof("|%s| sending one-off message to %s : '%s'", t.Type(), p.ToPhoneNumber, p.Content)

	return nil // TODO rm this block before deploying

	// reach out to AWS pinpoint and blast off an SMS
	resp, err := md.pp.SendMessages(ctx, &pinpoint.SendMessagesInput{
		ApplicationId: utils.Ptr("ecea11cc234a4af78bfe9831beca48bf"), // TODO pop this in ENV; its our Pinpoint APP ID
		MessageRequest: &ppt.MessageRequest{
			Addresses: map[string]ppt.AddressConfiguration{
				p.ToPhoneNumber: {ChannelType: ppt.ChannelTypeSms},
			},
			MessageConfiguration: &ppt.DirectMessageConfiguration{
				SMSMessage: &ppt.SMSMessage{
					// OriginationNumber: utils.Ptr("xxx"), // TODO different user different number
					Body:        &p.Content,
					MessageType: ppt.MessageTypePromotional,
				},
			},
		},
	})

	if err != nil {
		logger.Infof("|%s|error sending message", t.Type())
		return err
	}

	// we can tear this apart for cool metric
	// more importantly (TODO) check and handle Status (ie. 200? 400? 500?)
	logger.Info(utils.JSONDump(resp.MessageResponse))
	return nil
}
