package tests

import (
	"testing"

	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/japhy-tech/backend-test/internal/repository"
	"github.com/japhy-tech/backend-test/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreatePetUsecase(t *testing.T) {
	mockRepo := new(repository.MockPetRepository)
	usecase := usecase.NewPetUsecase(mockRepo)

	pet := &entity.CreatePet{
		Species:                  "dog",
		PetSize:                  "tall",
		Name:                     "doggo",
		AverageMaleAdultWeight:   60000,
		AverageFemaleAdultWeight: 58000,
	}

	createdPet := &entity.Pet{
		ID:                       1,
		Species:                  "dog",
		PetSize:                  "tall",
		Name:                     "doggo",
		AverageMaleAdultWeight:   60000,
		AverageFemaleAdultWeight: 58000,
	}

	mockRepo.On("Create", pet).Return(1, nil)

	result, err := usecase.CreatePet(pet)

	assert.NoError(t, err)
	assert.Equal(t, createdPet, result)
	mockRepo.AssertExpectations(t)
}
