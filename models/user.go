package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	FirstName string `json:"firstname" gorm:"serializer:json;column:firstname"`
	LastName  string `json:"lastname" gorm:"serializer:json;column:lastname"`
	Email     string `json:"email" gorm:"serializer:json;column:email"`
	ID        string `json:"id" gorm:"serializer:json;column:id"`
	Password  string `json:"password" gorm:"serializer:json;column:password"`
	CreatedAt string `json:"created_at" gorm:"serializer:json;column:created_at"`
}

// basic crud
func (u *User) Table() string {
	return "dbo.example"
}

func (u *User) ValidateEmail(db *gorm.DB, user User) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if user.Email != "" {
			return db
		}
		db.AddError(errors.New("invalid email provided"))
		return db
	}
}

func (u *User) Create(db *gorm.DB) error {
	if err := db.Table(u.Table()).Scopes(u.ValidateEmail(db, *u)).Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Update(db *gorm.DB, user User) error {
	return db.Table(u.Table()).Updates(&u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Table(u.Table()).Unscoped().Delete(&u).Error
}

func (u *User) Get(db *gorm.DB) *gorm.DB {
	return db.Table(u.Table())
}

func (u *User) GetByAttr(db *gorm.DB) *gorm.DB {
	return db.Table(u.Table()).First(u)
}

func (u *User) GetByIds(ids []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(u.Table()).Where("id IN (?)", ids).First(u)
	}
}
