package repository

import (
	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockPetRepository struct {
	mock.Mock
}

func (m *MockPetRepository) Create(pet *entity.CreatePet) (int, error) {
	args := m.Called(pet)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockPetRepository) GetAll() ([]entity.Pet, error) {
	args := m.Called()
	return args.Get(0).([]entity.Pet), args.Error(1)
}

func (m *MockPetRepository) GetByID(id int) (*entity.Pet, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Pet), args.Error(1)
}

func (m *MockPetRepository) Update(id int, pet *entity.UpdatePet) (int, error) {
	args := m.Called(id, pet)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockPetRepository) Delete(id int) (int, error) {
	args := m.Called(id)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockPetRepository) SearchPets(searchPets *entity.SearchPets) ([]entity.Pet, error) {
	args := m.Called(searchPets)
	return args.Get(0).([]entity.Pet), args.Error(1)
}
