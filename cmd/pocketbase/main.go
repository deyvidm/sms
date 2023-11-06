package main

import (
	"log"

	"github.com/deyvidm/sms/cmd/pocketbase/eventhooks"
	"github.com/deyvidm/sms/cmd/pocketbase/routes"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()
	ehs, err := eventhooks.NewStore(app)
	if err != nil {
		log.Fatal(err)
	}
	ehs.ApplyHooks()
	routes.RegisterCustomRoutes(app)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
