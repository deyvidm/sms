package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/deyvidm/sms/utils"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// This comes in Straight from SNS
type SNSEvent struct {
	Type             string `json:"Type"`
	Message          string `json:"Message"` // this is all stringified+escaped JSON => SNSRawMessage
	MessageId        string `json:"MessageId"`
	TopicArn         string `json:"TopicArn"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}

// This is SNSEvent.Message in struct form
type SNSRawMesage struct {
	OriginationNumber          string `json:"OriginationNumber"`
	DestinationNumber          string `json:"DestinationNumber"`
	MessageKeyword             string `json:"MessageKeyword"`
	MessageBody                string `json:"MessageBody"`
	PreviousPublishedMessageId string `json:"PreviousPublishedMessageId"`
	InboundMessageId           string `json:"InboundMessageId"`
}

func ListenForSNS(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/respond",
			Handler: func(c echo.Context) error {

				body, err := ioutil.ReadAll(c.Request().Body)
				if err != nil {
					log.Printf("Error reading body: %v", err)
				}
				notif := SNSEvent{}
				err = json.Unmarshal(body, &notif)
				if err != nil {
					log.Printf("Error unmarshalling SNS Notification: %v", err)
				}

				msg := SNSRawMesage{}
				err = json.Unmarshal([]byte(notif.Message), &msg)
				if err != nil {
					log.Printf("Error unmarshalling JSON: %v", err)
				}

				utils.JSONDump(msg)
				//
				// app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
				// 	// find corresponding sender contact record
				// 	records, err := txDao.FindRecordsByExpr("contact", dbx.HashExp{"phone": msg.OriginationNumber})
				// 	if err != nil {
				// 		log.Fatalln(err)
				// 		return err
				// 	}
				// 	if len(records) != 1 {
				// 		log.Println("found", len(records), "contact records for", msg.OriginationNumber)
				// 	}

				// 	records, err = txDao.FindRecordsByExpr("attendee", dbx.HashExp{"contact": records[0].Id, "event": eventID})
				// 	if err != nil {
				// 		log.Fatalln(err)
				// 		return err
				// 	}
				// 	if len(records) != 1 {
				// 		log.Println("found", len(records), "attendance for", phone)
				// 	}
				// 	for _, r := range records {
				// 		r.Set("status", status)
				// 		if err := txDao.SaveRecord(r); err != nil {
				// 			return err
				// 		}
				// 	}
				// 	return nil
				// })
				//
				return nil
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})
}
