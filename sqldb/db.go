package sqldb

import (
	"database/sql"
	"fmt"

	// Loads the mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the SQL connection to be used within the database package
func InitDB(dbUser string, dbPass string, dbAddr string, dbSchema string) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPass, dbAddr, dbSchema)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
