package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabase - Create a new database conection. Schema will be migrated if
// not found
func NewDatabase(username, password, host, port, database string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &HomeAddress{})

	return db
}
