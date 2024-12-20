package routes

import (
	"my-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterCharacterRoutes(router *mux.Router, controller *controllers.CharacterController) {
	router.HandleFunc("/Getcharacters", controller.GetAllCharacters).Methods(http.MethodGet)
	router.HandleFunc("/AddCharacter", controller.CreateCharacter).Methods(http.MethodPost)
}
