package main

import (
	"log"

	"github.com/b12o/pocket-docket/handler"
	"github.com/b12o/pocket-docket/middleware"
	"github.com/b12o/pocket-docket/util"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {

	app := pocketbase.New()

	// app needs to be injected into request contexts
	// in order for e.g DB operations to have access to the app object)
	// so include the app context using middleware
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(middleware.InjectPocketBaseAppMiddleware(app))
		return nil
	})

	//TODO: use middleware for authentication

	util.RegisterRoute(app, "GET", "/", handler.RootHandler)

	util.RegisterRoute(app, "GET", "/counter", handler.CountHandler)
	util.RegisterRoute(app, "POST", "/counter", handler.CountHandler)

	util.RegisterRoute(app, "POST", "/users", handler.HandleCreateUser)
	util.RegisterRoute(app, "GET", "/users/:userId", handler.HandleGetUser)
	util.RegisterRoute(app, "PATCH", "/users/:userId", handler.HandleUpdateUser)
	util.RegisterRoute(app, "DELETE", "/users/:userId", handler.HandleDeleteUser)

	util.RegisterRoute(app, "POST", "/tasks", handler.HandleCreateTask)
	util.RegisterRoute(app, "GET", "/tasks", handler.HandleGetTasks)
	util.RegisterRoute(app, "GET", "/tasks/:taskId", handler.HandleGetTask)
	util.RegisterRoute(app, "PATCH", "/tasks/:taskId", handler.HandleUpdateTask)
	util.RegisterRoute(app, "DELETE", "/tasks/:taskId", handler.HandleDeleteTask)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
