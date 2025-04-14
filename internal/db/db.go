package db

import (
	"database/sql"
	"os"
	"path"

	"github.com/CosmicPredator/chibi/internal"
	_ "modernc.org/sqlite"
)

type DbContext struct {
	dbConn *sql.DB
}

// creates required SQL tables
func (dc *DbContext) createRequiredTables() error {
	tx, err := dc.dbConn.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(QUERY_CREATE_TABLE); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// initialize an SQL instance.
// if isFirstTime is true, it'll clean the os config dir
// and creates a brand new table
func (dc *DbContext) init(isFirstTime bool) error {
	osConfigPath, _ := os.UserConfigDir()
	configDir := path.Join(osConfigPath, "chibi")

	if isFirstTime {
		internal.CreateConfigDir()
	}
	dbPath := path.Join(configDir, internal.DB_NAME)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	dc.dbConn = db
	if isFirstTime {
		err := dc.createRequiredTables()
		if err != nil {
			return err
		}
	}
	return nil
}

// add a key value pair to config table
func (dc *DbContext) SetConfig(key string, value string) error {
	tx, err := dc.dbConn.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(QUERY_INSERT_CONFIG, key, value); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// gets a key's value from the config table
func (dc *DbContext) GetConfig(key string) (*string, error) {
	tx, err := dc.dbConn.Begin()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(QUERY_GET_CONFIG, key)
	var value string

	if err = row.Scan(&value); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return &value, nil
}

// closes the DB instance after usage
func (dc *DbContext) Close() {
	dc.dbConn.Close()
}

// returns a brand new DbContext instance
func NewDbConn(isFirstTime bool) (*DbContext, error) {
	dbContext := DbContext{}
	err := dbContext.init(isFirstTime)
	if err != nil {
		return nil, err
	}
	return &dbContext, nil
}
