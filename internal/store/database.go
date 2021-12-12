package store

import (
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

// NewDatabase - Create a new database conection. Schema will be migrated if
// not found
func NewDatabase(username, password, host, port, database string) *sqlx.DB {
	connString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, database)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(100)
	runMigrations(db)
	return db
}

// NewTestDatabase - Create a new mock database conection
func NewTestDatabase() (*sqlx.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	mockDBSQLX := sqlx.NewDb(mockDB, "postgres") // returns *sqlx.DB
	return mockDBSQLX, mock, err
}

// runMigrations - Runs dataase migrations
func runMigrations(db *sqlx.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "../../schema",
	}
	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
