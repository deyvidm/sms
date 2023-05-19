package tasks

import (
	"encoding/json"
	"strings"

	"github.com/deyvidm/sms-asynq/types"
	"github.com/deyvidm/sms-asynq/utils"
	"github.com/hibiken/asynq"
)

const (
	TypeNewResponse = "response:new"
)

type NewResponsePayload struct {
	From    string
	Content string
}

func NewReponseTask(FromPhoneNumber, content string) (*asynq.Task, error) {
	payload, err := json.Marshal(NewResponsePayload{From: FromPhoneNumber, Content: content})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeNewResponse, payload), nil
}

var AcceptedResponseStrings = []string{"yes", "y"}
var DeclinedResponseStrings = []string{"no", "n"}

func (r *NewResponsePayload) parseYesNoReply(content string) (types.InviteStatus, error) {
	if utils.Contains(AcceptedResponseStrings, content) {
		return types.InviteStatus_Accepted, nil
	}
	if utils.Contains(DeclinedResponseStrings, content) {
		return types.InviteStatus_Declined, nil
	}
	return "", types.ResponseParseError{}
}

// this method parses user SMS responses and discovers the target invite,
// as well as what kind of response they would like to submit to the target invite
func (r *NewResponsePayload) Parse(parsedInfo *types.ResponseInfo) error {
	if parsedInfo == nil {
		parsedInfo = &types.ResponseInfo{}
	}
	terms := strings.Fields(strings.ToLower(strings.TrimSpace(r.Content)))

	for _, term := range terms {
		if utils.IsAllAlpha(term) {
			status, err := r.parseYesNoReply(term)
			if err != nil {
				return err
			}
			parsedInfo.Status = &status
		} else if utils.IsAllNumeric(term) {
			key, err := utils.ParseFloat64(term)
			if err != nil {
				return err
			}
			parsedInfo.TargetInviteKey = &key
		}
	}

	if parsedInfo.Status == nil {
		return types.ResponseParseError{}
	}

	return nil
}
