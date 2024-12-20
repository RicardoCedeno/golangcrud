package services

import (
	"my-api/interfaces"
	"my-api/models"
)

type CharacterService struct {
	repository interfaces.ICharacterRepository
}

func NewUserService(repo interfaces.ICharacterRepository) *CharacterService {
	return &CharacterService{repo}
}

func (s *CharacterService) GetAllCharacters() ([]models.Character, error) {
	return s.repository.GetAllCharacters()
}

func (s *CharacterService) CreateCharacter(character models.Character) error {
	return s.repository.CreateCharacter(character)
}
