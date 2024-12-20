package Repositories

import (
	"database/sql"
	"my-api/interfaces"
	"my-api/models"
)

type CharacterRepository struct {
	db *sql.DB
}

func NewChracterRepository(db *sql.DB) interfaces.ICharacterRepository {
	return &CharacterRepository{db}
}

func (r *CharacterRepository) GetAllCharacters() ([]models.Character, error) {
	rows, err := r.db.Query("SELECT * FROM dofus")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []models.Character
	for rows.Next() {
		var character models.Character
		if err := rows.Scan(&character.Id, &character.Name, &character.Class); err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}
	return characters, nil
}

func (r *CharacterRepository) CreateCharacter(character models.Character) error {
	_, err := r.db.Exec("INSERT INTO dofus (id, name, class) VALUES ($1, $2, $3)", character.Id, character.Name, character.Class)
	return err
}
