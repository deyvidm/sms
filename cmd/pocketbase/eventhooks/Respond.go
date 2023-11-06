package eventhooks

import (
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

// this event hook fires off every time a new response record is inserted
// i.e. on every successful POST to /api/respond
// i.e. every time an attendee responds to the event organizer
func (ehs *EventhookStore) Eventhook_OnCreateResponse() {
	ehs.app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "response" {
			phone := e.Record.GetString("originationNumber")
			records, err := ehs.app.Dao().FindRecordsByExpr("contact", dbx.HashExp{"phone": phone})
			if err != nil {
				log.Fatalf(err.Error())
			}
			if len(records) != 1 {
				log.Println("found", len(records), "contact records for", phone)
			}
			e.Record.Set("sender", records[0].Id)
			ehs.saveRecord(e.Record)
		}
		return nil
	})
}
