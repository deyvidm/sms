package workers

import (
	"context"
	"encoding/json"

	"github.com/deyvidm/sms/dispatcher/client"
	"github.com/deyvidm/sms/dispatcher/tasks"
	"github.com/deyvidm/sms/dispatcher/types"
	"github.com/deyvidm/sms/dispatcher/utils"
	"github.com/hibiken/asynq"
)

type ResponseProcessor struct {
	wbc *client.WebBackendClient
	irs *client.InviteResponseStore
}

func NewResponseProcessor(wbc *client.WebBackendClient, irs *client.InviteResponseStore) *ResponseProcessor {
	return &ResponseProcessor{
		wbc: wbc,
		irs: irs,
	}
}

func (rp *ResponseProcessor) HandleResponse(ctx context.Context, t *asynq.Task) error {
	var p tasks.NewResponsePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	logger.Infof("Response from %s : '%s'", p.From, p.Content)

	parsedInfo := types.ResponseInfo{}
	err := p.Parse(&parsedInfo)
	switch err {
	case types.ResponseParseError{}:
		logger.Errorf("bad response content: |%s|; prompt user to resubmit", p.Content)
		// TODO: create and queue new task prompting user for another response
		return nil
	case nil:
		break
	default:
		logger.Errorf("error while processing task: ", err)
		return err
	}

	targetInviteID, err := rp.irs.FetchTargetInviteID(p.From, parsedInfo)
	switch err {
	case types.MissingKeyError{}:
		logger.Errorf(err.Error())
		// TODO: create and queue new task prompting user for another response
		return nil
	case nil:
		break
	default:
		logger.Errorf("error while processing task: %s", err)
		return err
	}
	logger.Info("%+v", parsedInfo)
	if err = rp.wbc.UpdateInvite(client.UpdateInvite{
		ID:     targetInviteID,
		Status: utils.Ptr(parsedInfo.Status.String()),
	}); err != nil {
		return err
	}

	return rp.irs.PopInvite(p.From, targetInviteID)
}
