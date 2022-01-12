package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var Db *sql.DB

func Connect(connectionString string) error {
	var err error
	Db, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = Db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	fmt.Printf("Connected!\n")
	return nil
}

var Context = context.Background()
