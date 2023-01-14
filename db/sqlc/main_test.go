package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:G8keeper@localhost:5432/simple-bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	// Connect to db
	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}