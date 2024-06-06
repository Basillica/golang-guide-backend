package datamanager

import (
	"errors"

	"github.com/Basillica/golang-guide/service"
	"github.com/Basillica/golang-guide/userroles"
	"gorm.io/gorm"
)

type userQueryManager struct {
	table         string
	userRoleModel service.UserRoleModel
	db            *gorm.DB
	userRole      string
	userScopes    []string
}

func NewUserQueryManager(db *gorm.DB, userscopes []string, userrole string) (userQueryManager, error) {
	if db == nil || userrole == "" || userscopes == nil {
		return userQueryManager{}, errors.New("incomplete vars provided")
	}
	var role userroles.Role
	r := role.Role(userrole)
	model, err := r.RoleModel(userscopes)
	if err != nil {
		return userQueryManager{}, err
	}
	return userQueryManager{
		table:         "dbo.example",
		userRoleModel: model,
		userRole:      userrole,
		userScopes:    userscopes,
		db:            db,
	}, nil
}

func (u *userQueryManager) Get(db *gorm.DB) *gorm.DB {
	return db.Table(u.table).Scopes(u.userRoleModel.ParseUserScopes(u.db, u.userRoleModel.UserPermissionRead())...)
}
