package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	ID        string `json:"id" gorm:"primaryKey;column:user_id;serializer:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// basic crud
func (u *User) Table() string {
	return "dbo.users"
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

func (u *User) ValidateEmail2(db *gorm.DB, user User) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if user.Email != "" {
			return db
		}
		db.AddError(errors.New("invalid email provided"))
		return db
	}
}

func (u *User) Create(db *gorm.DB) error {
	if err := db.Table(u.Table()).Scopes(u.ValidateEmail(db, *u), u.ValidateEmail2(db, *u)).Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Update()       {}
func (u *User) Delete()       {}
func (u *User) Get(id string) {}

func (u *User) GetById() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(u.Table()).Where("id = (?)", u.ID).First(u)
	}
}

func (u *User) GetByIds(ids []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(u.Table()).Where("id IN (?)", ids).First(u)
	}
}
