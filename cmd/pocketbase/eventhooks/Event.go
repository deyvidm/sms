package eventhooks

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/deyvidm/sms/utils"
	"github.com/pocketbase/pocketbase/core"
)

func (ehs *EventhookStore) Eventhook_OnCreateEventSendSMS() {
	ehs.app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "attendee" {
			contact, err := ehs.app.Dao().FindRecordById("contact", e.Record.GetString("contact"))
			if err != nil {
				log.Fatalf(err.Error())
			}
			event, err := ehs.app.Dao().FindRecordById("event", e.Record.GetString("event"))
			if err != nil {
				log.Fatalf(err.Error())
			}
			phone := contact.GetString("phone")
			smsBody := event.GetString("description")
			log.Printf("inviting %s %s %s to %s:", contact.GetString("first_name"), contact.GetString("last_name"), contact.GetString("phone"), event.GetString("title"))
			client := sns.NewFromConfig(ehs.awscfg)
			resp, err := client.Publish(context.Background(), &sns.PublishInput{
				Message:     utils.Ptr(smsBody),
				PhoneNumber: utils.Ptr(phone),
			})
			if err != nil {
				log.Fatalln(err)
				return err
			}
			utils.JSONDump(resp)
			// TODO make status an enum; pocketbase native models?
			ehs.updateAttendanceStatus(event.Id, phone, "invited")
			if err != nil {
				log.Fatalln(err)
				return err
			}
		}
		return nil
	})
}
