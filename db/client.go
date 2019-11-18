package db

import (
	"io/ioutil"
	"sync"

	// Need this import for its side effects with sqlx.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

const (
	dsn    = "app:password@tcp(localhost:3306)/db?parseTime=true&clientFoundRows=true"
	driver = "mysql"
)

var (
	dbClient     *sqlx.DB
	dbClientSync sync.Once
)

// Client returns an initialised DB client.
// It will initialise the client and create the tables in the
// db/tables directory on first call.
func Client() *sqlx.DB {
	dbClientSync.Do(func() {
		client, err := sqlx.Connect(driver, dsn)
		if err != nil {
			panic(err)
		}

		client.Mapper = reflectx.NewMapper("json")

		err = initTables(client)
		if err != nil {
			panic(err)
		}

		dbClient = client
	})

	return dbClient
}

func initTables(client *sqlx.DB) error {
	const tablesDir = "./db/tables/"
	files, err := ioutil.ReadDir(tablesDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		contents, err := ioutil.ReadFile(tablesDir + file.Name())
		if err != nil {
			return err
		}

		_, err = client.Exec(string(contents))
		if err != nil {
			return err
		}
	}

	return nil
}
