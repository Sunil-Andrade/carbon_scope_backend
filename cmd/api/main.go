package main

import (
	"carbon/internal/config"
	"carbon/internal/database"
	"carbon/internal/router"
	"log"
	"net/http"
)

func main() {

	cfg := config.Load()

	database.Connect(cfg)

	r := router.SetupRouter()

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
