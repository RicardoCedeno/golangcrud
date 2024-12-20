package main

import (
	"log"
	"my-api/Repositories"
	"my-api/config"
	"my-api/controllers"
	"my-api/routes"
	"my-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	characterRepo := Repositories.NewChracterRepository(db)
	characterService := services.NewUserService(characterRepo)
	characterController := controllers.NewCharacterController(characterService)

	router := mux.NewRouter()
	routes.RegisterCharacterRoutes(router, characterController)

	log.Println("servidor iniciado en puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
