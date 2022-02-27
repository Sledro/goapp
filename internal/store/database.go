package store

import (
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

// NewDatabase - Create a new database conection. Schema will be migrated if
// not found
func NewDatabase(username, password, host, port, database string) *sqlx.DB {
	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, host, port, database)
	db, err := sqlx.Connect("mysql", connString)
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
	mockDBSQLX := sqlx.NewDb(mockDB, "mysql") // returns *sqlx.DB
	return mockDBSQLX, mock, err
}

// runMigrations - Runs dataase migrations
func runMigrations(db *sqlx.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "../../schema",
	}
	n, err := migrate.Exec(db.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
