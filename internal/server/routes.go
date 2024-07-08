package server

import (
	"database/sql"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/internal/delivery/http"
	"github.com/japhy-tech/backend-test/internal/repository"
	"github.com/japhy-tech/backend-test/internal/usecase"
)

type App struct {
	logger *charmLog.Logger
	db     *sql.DB
}

func NewApp(logger *charmLog.Logger, db *sql.DB) *App {
	return &App{
		logger: logger,
		db:     db,
	}
}

// TODO: améliorer cette partie
func (a *App) RegisterRoutes(r *mux.Router) {
	petRepo := repository.NewPetRepository(a.db)
	petUsecase := usecase.NewPetUsecase(petRepo)
	http.NewPetHandler(r, petUsecase, a.logger)
}
