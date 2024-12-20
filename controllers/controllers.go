package controllers

import (
	"encoding/json"
	"my-api/models"
	"my-api/services"
	"net/http"
)

type CharacterController struct {
	service *services.CharacterService
}

func NewCharacterController(service *services.CharacterService) *CharacterController {
	return &CharacterController{service}
}

func (c *CharacterController) GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := c.service.GetAllCharacters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(characters)
}

func (c *CharacterController) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var character models.Character
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := c.service.CreateCharacter(character); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(character)
}
