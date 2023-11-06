package eventhooks

import (
	"context"
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/pocketbase/pocketbase"
)

type EventhookStore struct {
	app    *pocketbase.PocketBase
	awscfg aws.Config
}

func NewStore(app *pocketbase.PocketBase) (*EventhookStore, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return &EventhookStore{app: app, awscfg: cfg}, nil
}

// now we can write new event hooks and have them be auto-applied
// func (*EventhookStore) EventHookt_AfterCreateResponeResponse() {}

func (e *EventhookStore) ApplyHooks() {
	// reflect and execute only the methods
	// that start with "EventHook_"
	// Get reflect.Value of the struct variable
	ehstoreRType := reflect.TypeOf(e)
	ehstoreRVal := reflect.ValueOf(e)
	// Loop through all methods of the struct variable
	for i := 0; i < ehstoreRType.NumMethod(); i++ {
		method := ehstoreRType.Method(i)
		name := method.Name
		// Check if method name starts with "EventHook_"
		if len(name) > 10 && name[:10] == "Eventhook_" {
			// Call the method
			ehstoreRVal.MethodByName(name).Call(nil)
			log.Println("Mounted ", name)
		}
	}
}
