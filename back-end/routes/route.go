package route

import (
	"net/http"

	"fullstack-project/back-end/controller"
	"fullstack-project/back-end/middleware"
)

func RegisterRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	// Authentication
	mux.HandleFunc("POST /register", controller.Register)
	mux.HandleFunc("POST /login", controller.Login)

	// User
	mux.Handle(
		"GET /users",
		middleware.Logger(
			middleware.JWT(
				middleware.AdminOnly(
					http.HandlerFunc(controller.GetUsers),
				),
			),
		),
	)

	mux.Handle(
		"GET /users/{id}",
		middleware.Logger(
			middleware.JWT(
				http.HandlerFunc(controller.GetUser),
			),
		),
	)

	mux.Handle(
		"POST /users",
		middleware.Logger(
			middleware.JWT(
				middleware.AdminOnly(
					http.HandlerFunc(controller.CreateUser),
				),
			),
		),
	)

	mux.Handle(
		"PUT /users/{id}",
		middleware.Logger(
			middleware.JWT(
				middleware.AdminOnly(
					http.HandlerFunc(controller.UpdateUser),
				),
			),
		),
	)

	mux.Handle(
		"DELETE /users/{id}",
		middleware.Logger(
			middleware.JWT(
				middleware.AdminOnly(
					http.HandlerFunc(controller.DeleteUser),
				),
			),
		),
	)

	// Job
	mux.Handle(
		"POST /jobs",
		middleware.Logger(
			middleware.JWT(
				http.HandlerFunc(controller.CreateScrapJob),
			),
		),
	)

	mux.Handle(
		"GET /jobs",
		middleware.Logger(
			middleware.JWT(
				http.HandlerFunc(controller.GetJobs),
			),
		),
	)

	mux.Handle(
		"DELETE /jobs/{id}",
		middleware.Logger(
			middleware.JWT(
				http.HandlerFunc(controller.DeleteJob),
			),
		),
	)

	// History
	mux.Handle(
		"GET /history",
		middleware.Logger(
			middleware.JWT(
				http.HandlerFunc(controller.GetHistory),
			),
		),
	)

	mux.Handle(
		"GET /history/{id}",
		middleware.Logger(
			middleware.JWT(
				http.HandlerFunc(controller.GetHistoryByID),
			),
		),
	)

	// Automation
	mux.Handle(
		"POST /shorts",
		middleware.Logger(
			middleware.JWT(
				http.HandlerFunc(controller.ReceiveShort),
			),
		),
	)

	return mux
}