package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// ---------- UTILS -----------

func main() {

	app := pocketbase.New()

	// app needs to be injected into request contexts
	// in order for e.g DB operations to have access to the app object)
	// so include the app context using middleware
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(InjectPocketBaseAppMiddleWare(app))
		return nil
	})

	RegisterRoute(app, "GET", "/", RootHandler)
	RegisterRoute(app, "GET", "/counter", CountHandler)
	RegisterRoute(app, "POST", "/counter", CountHandler)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
