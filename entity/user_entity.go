package entity

import (
	"hacktiv8_fp_2/helpers"

	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Age      int    `json:"age"`
	BaseModel
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashAndSalt(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashAndSalt(u.Password)
	if err != nil {
		return err
	}
	return nil
}
