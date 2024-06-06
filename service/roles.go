package service

import "gorm.io/gorm"

type UserRoleModel interface {
	UserPermissionRead() func(db *gorm.DB) *gorm.DB
	UserPermissionWrite() func(db *gorm.DB) *gorm.DB
	//
	ParseUserScopes(db *gorm.DB, baseScopes ...func(db *gorm.DB) *gorm.DB) []func(db *gorm.DB) *gorm.DB
}
