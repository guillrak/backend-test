package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	charmLog "github.com/charmbracelet/log"
)

var (
	once       sync.Once
	dbInstance *sql.DB
)

func NewMysqlDB(logger *charmLog.Logger) *sql.DB {
	once.Do(func() {
		mysqlDSN := fmt.Sprintf("root:%s@(mysql-test:3306)/core?parseTime=true", os.Getenv("MYSQL_ROOT_PASSWORD"))

		db, err := sql.Open("mysql", mysqlDSN)
		if err != nil {
			logger.Fatal(err.Error())
		}

		dbInstance = db
	})

	return dbInstance
}

func GetDb() *sql.DB {
	return dbInstance
}
