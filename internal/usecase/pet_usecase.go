package usecase

import (
	"fmt"

	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/japhy-tech/backend-test/internal/repository"
)

type PetUsecase interface {
	CreatePet(pet *entity.CreatePet) (*entity.Pet, error)
	GetPets() ([]entity.Pet, error)
	GetPetByID(id int) (*entity.Pet, error)
	UpdatePet(id int, pet *entity.UpdatePet) (*entity.Pet, error)
	DeletePet(id int) error
	SearchPets(searchPets *entity.SearchPets) ([]entity.Pet, error)
}

type petUsecase struct {
	petRepo repository.PetRepository
}

func NewPetUsecase(petRepo repository.PetRepository) PetUsecase {
	return &petUsecase{petRepo: petRepo}
}

func (u *petUsecase) CreatePet(pet *entity.CreatePet) (*entity.Pet, error) {
	id, err := u.petRepo.Create(pet)
	if err != nil {
		return nil, err
	}

	createdPet := &entity.Pet{
		ID:                       id,
		Species:                  pet.Species,
		PetSize:                  pet.PetSize,
		Name:                     pet.Name,
		AverageMaleAdultWeight:   pet.AverageMaleAdultWeight,
		AverageFemaleAdultWeight: pet.AverageFemaleAdultWeight,
	}

	return createdPet, nil
}

func (u *petUsecase) GetPets() ([]entity.Pet, error) {
	return u.petRepo.GetAll()
}

func (u *petUsecase) GetPetByID(id int) (*entity.Pet, error) {
	return u.petRepo.GetByID(id)
}

func (u *petUsecase) UpdatePet(id int, pet *entity.UpdatePet) (*entity.Pet, error) {
	rowsAffected, err := u.petRepo.Update(id, pet)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("ID not found")
	}

	updatedPet := &entity.Pet{
		ID:                       id,
		Species:                  pet.Species,
		PetSize:                  pet.PetSize,
		Name:                     pet.Name,
		AverageMaleAdultWeight:   pet.AverageMaleAdultWeight,
		AverageFemaleAdultWeight: pet.AverageFemaleAdultWeight,
	}

	return updatedPet, nil
}

func (u *petUsecase) DeletePet(id int) error {
	rowsAffected, err := u.petRepo.Delete(id)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ID not found")
	}

	return nil
}

func (u *petUsecase) SearchPets(searchPets *entity.SearchPets) ([]entity.Pet, error) {
	return u.petRepo.SearchPets(searchPets)
}
