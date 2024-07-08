package database_actions

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// FillPetsTable fill the "pets" table with the contents of the CSV file.
//
// The table will only be filled if it is empty.
//
// Returns the number of rows affected or an error.
func LoadPetsTable(db *sql.DB, filePath string) (int64, error) {

	//Check if the table is empty
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM pets").Scan(&count)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		mysql.RegisterLocalFile(filePath)

		query := fmt.Sprintf(
			`LOAD DATA LOCAL INFILE '%s' 
			INTO TABLE pets 
			FIELDS TERMINATED BY ","
			LINES TERMINATED BY "\n"
			IGNORE 1 LINES
			`, filePath)

		result, err := db.Exec(query)
		if err != nil {
			return 0, err
		}

		return result.RowsAffected()
	}

	return 0, nil
}
