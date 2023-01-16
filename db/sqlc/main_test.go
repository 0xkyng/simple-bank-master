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
	// Create a new connection 
	// To the db
	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	// Use conn to create new testQueries
	testQueries = New(conn)

	os.Exit(m.Run())
}