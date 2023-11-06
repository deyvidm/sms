package eventhooks

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
)

func (ehs *EventhookStore) Eventhook_CreateContactInAWS() {
	ehs.app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "contact" {
			phone := e.Record.GetString("phone")
			client := sns.NewFromConfig(ehs.awscfg)
			log.Printf("sending [%s] to aws:", phone)
			_, err := client.CreateSMSSandboxPhoneNumber(context.TODO(),
				&sns.CreateSMSSandboxPhoneNumberInput{
					PhoneNumber: &phone,
				})
			if err != nil {
				log.Fatalln(err)
				return err
			}
			err = ehs.updateContactVerifiedStatus(phone, string(types.SMSSandboxPhoneNumberVerificationStatusPending))
			if err != nil {
				log.Fatalln(err)
				return err
			}
		}
		return nil
	})
}

func (ehs *EventhookStore) updateContactVerifiedStatus(phone string, status string) error {
	return ehs.app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
		records, err := txDao.FindRecordsByExpr("contact", dbx.HashExp{"phone": phone})
		if err != nil {
			log.Fatalln(err)
			return err
		}
		if len(records) != 1 {
			log.Println("found", len(records), "for", phone)
		}
		for _, r := range records {
			r.Set("status", status)
			if err := txDao.SaveRecord(r); err != nil {
				return err
			}
		}
		return nil
	})
}

func (ehs *EventhookStore) updateAttendanceStatus(eventID, phone, status string) error {
	return ehs.app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
		records, err := txDao.FindRecordsByExpr("contact", dbx.HashExp{"phone": phone})
		if err != nil {
			log.Fatalln(err)
			return err
		}
		if len(records) != 1 {
			log.Println("found", len(records), "contact records for", phone)
		}

		records, err = txDao.FindRecordsByExpr("attendee", dbx.HashExp{"contact": records[0].Id, "event": eventID})
		if err != nil {
			log.Fatalln(err)
			return err
		}
		if len(records) != 1 {
			log.Println("found", len(records), "attendance for", phone)
		}
		for _, r := range records {
			r.Set("status", status)
			if err := txDao.SaveRecord(r); err != nil {
				return err
			}
		}
		return nil
	})
}

func (ehs *EventhookStore) Eventhook_ListContacts() {
	ehs.app.OnRecordsListRequest().Add(func(e *core.RecordsListEvent) error {
		if e.Collection.Name == "contact" {
			client := sns.NewFromConfig(ehs.awscfg)
			log.Printf("fetching & updating AWS contacts")
			resp, err := client.ListSMSSandboxPhoneNumbers(context.TODO(), &sns.ListSMSSandboxPhoneNumbersInput{})
			if err != nil {
				log.Fatalf(err.Error())
			}

			// this whole block just updates the verification status in our DB.
			// this code should ideally run in a separate service,
			// dedicated to monitoring the database for changes and periodically syncing with AWS
			for _, p := range resp.PhoneNumbers {
				ehs.updateContactVerifiedStatus(*p.PhoneNumber, string(p.Status))
				if err != nil {
					log.Fatalln(err)
					return err
				}
			}

		}
		return nil
	})
}
