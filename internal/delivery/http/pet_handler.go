package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/internal/entity"
	"github.com/japhy-tech/backend-test/internal/usecase"
)

type PetHandler struct {
	PetUsecase usecase.PetUsecase
	logger     *charmLog.Logger
}

func NewPetHandler(router *mux.Router, pu usecase.PetUsecase, logger *charmLog.Logger) {
	handler := &PetHandler{
		PetUsecase: pu,
		logger:     logger,
	}

	router.HandleFunc("/pets", handler.CreatePet).Methods("POST")
	router.HandleFunc("/pets", handler.GetPets).Methods("GET")
	router.HandleFunc("/pets/{id}", handler.GetPet).Methods("GET")
	router.HandleFunc("/pets/{id}", handler.UpdatePet).Methods("PUT")
	router.HandleFunc("/pets/{id}", handler.DeletePet).Methods("DELETE")
	router.HandleFunc("/pets/search", handler.SearchPets).Methods("POST")
}

// CreatePet godoc
// @Summary Create a new pet
// @Description Adds a new pet to the database
// @Tags Pet
// @Accept json
// @Produce json
// @Param CreatePet body entity.CreatePet true "Pet object"
// @Success 201 {object} SuccessResponse{data=entity.Pet}
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /v1/pets [post]
func (h *PetHandler) CreatePet(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("[POST]	/v1/pets")

	var pet entity.CreatePet

	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		SendError(w, http.StatusBadRequest, err.Error())
		h.logger.Error("[POST]	/v1/pets; error:", err.Error())
		return
	}

	createdPet, err := h.PetUsecase.CreatePet(&pet)
	if err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		h.logger.Error("[POST]	/v1/pets; error:", err.Error())
		return
	}

	SendSuccess(w, http.StatusOK, createdPet)
}

// GetPets godoc
// @Summary Get all pets
// @Description Get all pets from the database
// @Tags Pet
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse{data=[]entity.Pet}
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /v1/pets [get]
func (h *PetHandler) GetPets(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("[GET]	/v1/pets")

	pets, err := h.PetUsecase.GetPets()
	if err != nil {
		SendError(w, http.StatusBadRequest, err.Error())
		h.logger.Error("[GET]	/v1/pets; error:", err.Error())
		return
	}

	SendSuccess(w, http.StatusOK, pets)
}

// GetPet godoc
// @Summary Get a pet
// @Description Get a pet by its ID
// @Tags Pet
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} SuccessResponse{data=entity.Pet}
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /v1/pets/{id} [get]
func (h *PetHandler) GetPet(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("[GET]	/v1/pets/{id}")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendError(w, http.StatusBadRequest, "Invalid ID")
		h.logger.Error("[GET]	/v1/pets/{id}; error:", err.Error())
		return
	}

	pet, err := h.PetUsecase.GetPetByID(id)
	if err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		h.logger.Error("[GET]	/v1/pets/{id}; error:", err.Error())
		return
	}

	SendSuccess(w, http.StatusOK, pet)
}

// UpdatePet godoc
// @Summary Update an existing pet
// @Description Updates the details of an existing pet
// @Tags Pet
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Param UpdatePet body entity.UpdatePet true "Pet object"
// @Success 200 {object} SuccessResponse{data=entity.Pet}
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /v1/pets/{id} [put]
func (h *PetHandler) UpdatePet(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("[PUT]	/v1/pets/{id}")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendError(w, http.StatusBadRequest, "Invalid ID")
		h.logger.Error("[PUT]	/v1/pets/{id}; error:", err.Error())
		return
	}

	var pet entity.UpdatePet
	err = json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		SendError(w, http.StatusBadRequest, err.Error())
		h.logger.Error("[PUT]	/v1/pets/{id}; error:", err.Error())
		return
	}

	updatedPet, err := h.PetUsecase.UpdatePet(id, &pet)
	if err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		h.logger.Error("[PUT]	/v1/pets/{id}; error:", err.Error())
		return
	}

	SendSuccess(w, http.StatusOK, updatedPet)
}

// DeletePet godoc
// @Summary Delete a pet
// @Description Delete a pet by its ID
// @Tags Pet
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} SuccessResponse{data=nil}
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /v1/pets/{id} [delete]
func (h *PetHandler) DeletePet(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("[DELETE]	/v1/pets/{id}")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendError(w, http.StatusBadRequest, "Invalid ID")
		h.logger.Error("[DELETE]	/v1/pets/{id}; error:", err.Error())
		return
	}

	err = h.PetUsecase.DeletePet(id)
	if err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		h.logger.Error("[DELETE]	/v1/pets/{id}; error:", err.Error())
		return
	}

	SendSuccess(w, http.StatusOK, nil)
}

// SearchPets godoc
// @Summary Search pets
// @Description Search for pets by species and weight
// @Tags Pet
// @Accept json
// @Produce json
// @Param SearchPets body entity.SearchPets true "Search options"
// @Success 200 {object} SuccessResponse{data=[]entity.Pet}
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /v1/pets/search [post]
func (h *PetHandler) SearchPets(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("[POST]	/v1/pets/search")

	var searchPets entity.SearchPets
	err := json.NewDecoder(r.Body).Decode(&searchPets)
	if err != nil {
		SendError(w, http.StatusBadRequest, err.Error())
		h.logger.Error("[POST]	/v1/pets/search; error:", err.Error())
		return
	}

	pets, err := h.PetUsecase.SearchPets(&searchPets)
	if err != nil {
		SendError(w, http.StatusInternalServerError, err.Error())
		h.logger.Error("[POST]	/v1/pets/search; error:", err.Error())
		return
	}

	SendSuccess(w, http.StatusOK, pets)
}
