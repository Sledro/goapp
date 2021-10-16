package storage

import (
	"gorm.io/gorm"
)

// Login - Returns user of provided email
func Login(user User, db *gorm.DB) error {
	err := db.Debug().Model(User{}).Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		return err
	}
	return nil
}
