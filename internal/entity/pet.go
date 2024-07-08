package entity

type Pet struct {
	ID                       int    `json:"id"`
	Species                  string `json:"species"`
	PetSize                  string `json:"pet_size"`
	Name                     string `json:"name"`
	AverageMaleAdultWeight   uint   `json:"average_male_adult_weight"`
	AverageFemaleAdultWeight uint   `json:"average_female_adult_weight"`
}

type CreatePet struct {
	Species                  string `json:"species"`
	PetSize                  string `json:"pet_size"`
	Name                     string `json:"name"`
	AverageMaleAdultWeight   uint   `json:"average_male_adult_weight"`
	AverageFemaleAdultWeight uint   `json:"average_female_adult_weight"`
}

type UpdatePet struct {
	Species                  string `json:"species"`
	PetSize                  string `json:"pet_size"`
	Name                     string `json:"name"`
	AverageMaleAdultWeight   uint   `json:"average_male_adult_weight"`
	AverageFemaleAdultWeight uint   `json:"average_female_adult_weight"`
}

type SearchPets struct {
	Species   string `json:"species"`
	MinWeight uint   `json:"min_weight"`
	MaxWeight uint   `json:"max_weight"`
}
