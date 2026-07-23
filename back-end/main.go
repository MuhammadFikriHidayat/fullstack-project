package main

import (
	"fmt"
	"log"
	"fullstack-project/back-end/automation"
	"fullstack-project/back-end/config"
	"fullstack-project/back-end/routes"
	"net/http"
)

func main() {
	config.ConnectDatabase()
	// e := echo.New()

	router := route.RegisterRoutes()

	fmt.Println("Automation Backend Running")
	fmt.Println("Server : http://localhost:8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}

    http.HandleFunc("/shorts", automation.ShortsHandler)
    http.ListenAndServe(":8080", nil)
}