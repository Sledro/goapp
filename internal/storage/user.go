package storage

import (
	"gorm.io/gorm"
)

// User - An app user
type User struct {
	gorm.Model `json:"-"`
	ID         int         `json:"user_id" gorm:"primaryKey,autoIncrement"`
	Firstname  string      `json:"firstname" validate:"required,min=3,max=100"`
	Lastname   string      `json:"lastname" validate:"required,min=3,max=100"`
	Username   string      `json:"username" validate:"required,min=3,max=30"`
	Password   string      `gorm:"size:100" json:"password,omitempty" validate:"required,min=8,max=30"`
	Email      string      `json:"email" validate:"required,email"`
	Address    HomeAddress `json:"home_address" gorm:"foreignKey:UserID" validate:"required"`
}

// HomeAddress - Users home address
type HomeAddress struct {
	gorm.Model `json:"-"`
	UserID     int    `json:"user_id"`
	Street1    string `json:"street1" validate:"required,min=3,max=100"`
	Street2    string `json:"street2" validate:"min=0,max=100"`
	Town       string `json:"town" validate:"required,min=3,max=100"`
	City       string `json:"city" validate:"required,min=3,max=100"`
	Country    string `json:"country" validate:"required,min=3,max=100"`
}

// UserCreate - Creates a user
func UserCreate(user User, db *gorm.DB) error {
	db.Create(&user)
	return nil
}

// UserGet - Gets a user
func UserGet(user User, db *gorm.DB) (User, error) {
	err := db.Where(&user).First(&user).Error
	if err != nil {
		return User{}, err
	}
	address := HomeAddress{}
	db.First(&address, user.ID)
	user.Address = address
	return user, nil
}

// UserUpdate - Updates a user
func UserUpdate(user User, userNew User, db *gorm.DB) (User, error) {
	db.Model(&user).Updates(userNew)
	db.Save(&user)
	return user, nil
}

// UserDelete - Deletes a user
func UserDelete(user User, db *gorm.DB) error {
	err := db.Where(&user).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// UserList - Gets a list of users
func UserList(db *gorm.DB) ([]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}
