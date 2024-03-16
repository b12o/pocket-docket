package main

import (
	"log"

	"github.com/b12o/pocket-docket/handlers"
	"github.com/b12o/pocket-docket/middlewares"
	"github.com/b12o/pocket-docket/utils"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {

	app := pocketbase.New()

	// app needs to be injected into request contexts
	// in order for e.g DB operations to have access to the app object)
	// so include the app context using middleware
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(middlewares.InjectPocketBaseAppMiddleware(app))
		return nil
	})

	//TODO: use middleware for authentication

	utils.RegisterRoute(app, "GET", "/", handlers.RootHandler)

	utils.RegisterRoute(app, "GET", "/counter", handlers.CountHandler)
	utils.RegisterRoute(app, "POST", "/counter", handlers.CountHandler)

	utils.RegisterRoute(app, "POST", "/users", handlers.CreateUserHandler)
	utils.RegisterRoute(app, "GET", "/users/:userId", handlers.GetUserHandler)
	utils.RegisterRoute(app, "PATCH", "/users/:userId", handlers.UpdateUserHandler)
	utils.RegisterRoute(app, "DELETE", "/users/:userId", handlers.DeleteUserHandler)

	utils.RegisterRoute(app, "POST", "/tasks", handlers.CreateTaskHandler)
	utils.RegisterRoute(app, "GET", "/tasks/:taskId", handlers.GetTaskHandler)
	utils.RegisterRoute(app, "PATCH", "/tasks/:taskId", handlers.UpdateTaskHandler)
	utils.RegisterRoute(app, "DELETE", "/tasks/:taskId", handlers.DeleteTaskHandler)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
