package repository

import (
	"database/sql"

	"github.com/japhy-tech/backend-test/internal/entity"
)

type PetRepository interface {
	Create(pet *entity.CreatePet) (int, error)
	GetAll() ([]entity.Pet, error)
	GetByID(id int) (*entity.Pet, error)
	Update(id int, pet *entity.UpdatePet) (int, error)
	Delete(id int) (int, error)
	SearchPets(searchPets *entity.SearchPets) ([]entity.Pet, error)
}

type petRepository struct {
	DB *sql.DB
}

func NewPetRepository(db *sql.DB) PetRepository {
	return &petRepository{DB: db}
}

func (r *petRepository) Create(pet *entity.CreatePet) (int, error) {
	result, err := r.DB.Exec(`
		INSERT INTO pets (species, pet_size, name, average_male_adult_weight, average_female_adult_weight) 
		VALUES (?, ?, ?, ?, ?)
	`, pet.Species, pet.PetSize, pet.Name, pet.AverageMaleAdultWeight, pet.AverageFemaleAdultWeight)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *petRepository) GetAll() ([]entity.Pet, error) {
	rows, err := r.DB.Query("SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight FROM pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []entity.Pet
	for rows.Next() {
		var pet entity.Pet
		err := rows.Scan(&pet.ID, &pet.Species, &pet.PetSize, &pet.Name, &pet.AverageMaleAdultWeight, &pet.AverageFemaleAdultWeight)
		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func (r *petRepository) GetByID(id int) (*entity.Pet, error) {
	var pet entity.Pet

	err := r.DB.QueryRow(`
		SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight 
		FROM pets 
		WHERE id = ?
	`, id).Scan(&pet.ID, &pet.Species, &pet.PetSize, &pet.Name, &pet.AverageMaleAdultWeight, &pet.AverageFemaleAdultWeight)
	if err != nil {
		return nil, err
	}

	return &pet, nil
}

func (r *petRepository) Update(id int, pet *entity.UpdatePet) (int, error) {
	result, err := r.DB.Exec(`
		UPDATE pets 
		SET species = ?, pet_size = ?, name = ?, average_male_adult_weight = ?, average_female_adult_weight = ? 
		WHERE id = ?`, pet.Species, pet.PetSize, pet.Name, pet.AverageMaleAdultWeight, pet.AverageFemaleAdultWeight, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func (r *petRepository) Delete(id int) (int, error) {
	result, err := r.DB.Exec(`
		DELETE FROM pets 
		WHERE id = ?
	`, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func (r *petRepository) SearchPets(searchPets *entity.SearchPets) ([]entity.Pet, error) {
	query := `
		SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight 
		FROM pets 
		WHERE 1=1
	`

	var args []interface{}

	if searchPets.Species != "" {
		query += " AND species = ?"
		args = append(args, searchPets.Species)
	}

	if searchPets.MinWeight > 0 {
		query += " AND (average_male_adult_weight >= ? OR average_female_adult_weight >= ?)"
		args = append(args, searchPets.MinWeight, searchPets.MinWeight)
	}

	if searchPets.MaxWeight > 0 {
		query += " AND (average_male_adult_weight <= ? OR average_female_adult_weight <= ?)"
		args = append(args, searchPets.MaxWeight, searchPets.MaxWeight)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []entity.Pet
	for rows.Next() {
		var pet entity.Pet
		err := rows.Scan(&pet.ID, &pet.Species, &pet.PetSize, &pet.Name, &pet.AverageMaleAdultWeight, &pet.AverageFemaleAdultWeight)
		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}
