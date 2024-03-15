package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {

	app := pocketbase.New()

	// app needs to be injected into request contexts
	// in order for e.g DB operations to have access to the app object)
	// so include the app context using middleware
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(InjectPocketBaseAppMiddleware(app))
		return nil
	})

	//TODO: use middleware for authentication

	RegisterRoute(app, "GET", "/", RootHandler)

	RegisterRoute(app, "GET", "/counter", CountHandler)
	RegisterRoute(app, "POST", "/counter", CountHandler)

	RegisterRoute(app, "POST", "/users", CreateUserHandler)
	RegisterRoute(app, "GET", "/users/:userId", GetUserHandler)
	RegisterRoute(app, "PATCH", "/users/:userId", UpdateUserHandler)
	RegisterRoute(app, "DELETE", "/users/:userId", DeleteUserHandler)

	RegisterRoute(app, "POST", "/tasks", CreateTaskHandler)
	RegisterRoute(app, "GET", "/tasks/:taskId", GetTaskHandler)
	RegisterRoute(app, "PATCH", "/tasks/:taskId", UpdateTaskHandler)
	RegisterRoute(app, "DELETE", "/tasks/:taskId", DeleteTaskHandler)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
