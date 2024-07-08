package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/japhy-tech/backend-test/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestSearchPets(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewPetRepository(db)

	searchCriteria := &entity.SearchPets{
		Species:   "dog",
		MinWeight: 40,
		MaxWeight: 70,
	}

	rows := sqlmock.NewRows([]string{"id", "species", "pet_size", "name", "average_male_adult_weight", "average_female_adult_weight"}).
		AddRow(2, "dog", "small", "little_one", 60, 50)

	query := `
		SELECT id, species, pet_size, name, average_male_adult_weight, average_female_adult_weight 
		FROM pets 
		WHERE 1=1 AND species = \? AND \(average_male_adult_weight >= \? OR average_female_adult_weight >= \?\) AND \(average_male_adult_weight <= \? OR average_female_adult_weight <= \?\)
	`

	mock.ExpectQuery(query).
		WithArgs(searchCriteria.Species, searchCriteria.MinWeight, searchCriteria.MinWeight, searchCriteria.MaxWeight, searchCriteria.MaxWeight).
		WillReturnRows(rows)

	result, err := repo.SearchPets(searchCriteria)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "little_one", result[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}
