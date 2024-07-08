package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/database_actions"
	"github.com/japhy-tech/backend-test/internal/database"
	"github.com/japhy-tech/backend-test/internal/server"

	_ "github.com/japhy-tech/backend-test/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	ApiPort        = "5000"
	BreedsFilePath = "database_actions/seeds/breeds.csv"
)

func main() {
	logger := charmLog.NewWithOptions(os.Stderr, charmLog.Options{
		Formatter:       charmLog.TextFormatter,
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "🧑‍💻 backend-test",
		Level:           charmLog.DebugLevel,
	})

	err := validateEnv()
	if err != nil {
		logger.Fatal(err.Error())
	}

	db := database.NewMysqlDB(logger)

	defer db.Close()
	db.SetMaxIdleConns(0)

	err = db.Ping()
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Database connected")

	err = database_actions.InitMigrator(db)
	if err != nil {
		logger.Fatal(err.Error())
	}

	msg, err := database_actions.RunMigrate("up", 0)
	if err != nil {
		logger.Fatal(err.Error())
	} else {
		logger.Info(msg)
	}

	// Loading data into the pets table
	nbRowsAffected, err := database_actions.LoadPetsTable(db, BreedsFilePath)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Unable to load pets table %s", err.Error()))
	}
	if nbRowsAffected > 0 {
		logger.Info(fmt.Sprintf("%d lines were successfully loaded into the pets table", nbRowsAffected))
	}

	app := server.NewApp(logger, db)

	r := mux.NewRouter()
	app.RegisterRoutes(r.PathPrefix("/v1").Subrouter())

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	err = http.ListenAndServe(
		net.JoinHostPort("", ApiPort),
		r,
	)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Unable to start service %s", err.Error()))
	}

	// =============================== Starting Msg ===============================
	logger.Info(fmt.Sprintf("Service started and listen on port %s", ApiPort))
}

// validateEnv checks that all the environment variables required to run the app are set.
func validateEnv() error {
	if os.Getenv("MYSQL_ROOT_PASSWORD") == "" {
		return fmt.Errorf("MYSQL_ROOT_PASSWORD is not set")
	}

	return nil
}
