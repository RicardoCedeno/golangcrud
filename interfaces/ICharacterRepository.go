package interfaces

import (
	"my-api/models"
)

type ICharacterRepository interface {
	GetAllCharacters() ([]models.Character, error)
	CreateCharacter(character models.Character) error
}
