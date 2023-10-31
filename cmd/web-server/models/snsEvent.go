package models

import (
	"encoding/json"
	"log"

	"github.com/deyvidm/sms/cmd/web-server/types"
)

// This comes in Straight from SNS
type SNSEvent struct {
	BaseModel
	Type             string
	MessageRaw       string
	RawMessageRef    string
	Message          SNSRawMesage `gorm:"foreignKey:RawMessageRef"`
	MessageId        string
	TopicArn         string
	Timestamp        string
	SignatureVersion string
	Signature        string
	SigningCertURL   string
	UnsubscribeURL   string
}

// This is SNSEvent.Message in struct form
type SNSRawMesage struct {
	BaseModel
	OriginationNumber          string
	DestinationNumber          string
	MessageKeyword             string
	MessageBody                string
	PreviousPublishedMessageId string
	InboundMessageId           string
}

func SNSEventFromInput(input types.SNSEvent) (SNSEvent, error) {
	msg := SNSRawMesage{}
	if err := json.Unmarshal([]byte(input.Message), &msg); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
	}

	return SNSEvent{
		Message:          msg,
		MessageId:        input.MessageId,
		TopicArn:         input.TopicArn,
		Timestamp:        input.Timestamp,
		SignatureVersion: input.SignatureVersion,
		Signature:        input.Signature,
		SigningCertURL:   input.SigningCertURL,
		UnsubscribeURL:   input.UnsubscribeURL,
	}, nil
}

func (e *SNSEvent) Save() (*SNSEvent, error) {
	err := DB.Create(&e).Error
	return e, err
}
