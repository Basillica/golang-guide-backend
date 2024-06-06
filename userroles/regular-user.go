package userroles

import "gorm.io/gorm"

type RegularUser struct {
	Role       string   // GLOBAL_ADMIN
	UserScopes []string // scopes that define a user's permission
	Roles      []string // list of possible roles in the application
}

// org1
func NewRegularUser(r Role, scopes []string) *RegularUser {
	return &RegularUser{
		Role:       r.String(),
		Roles:      ROLES,
		UserScopes: scopes,
	}
}

func (ga *RegularUser) UserPermissionRead() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, scope := range ga.UserScopes {
			db.Where("scope LIKE ?", `%"`+scope+`%`)
		}
		return db
	}
}

func (ga *RegularUser) UserPermissionWrite() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, scope := range ga.UserScopes {
			db.Where("scope LIKE ?", `%"`+scope+`%`)
		}
		return db
	}
}

func (ga *RegularUser) ParseUserScopes(db *gorm.DB, baseScopes ...func(db *gorm.DB) *gorm.DB) []func(db *gorm.DB) *gorm.DB {
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, baseScopes...)
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db
	})
	return scopes
}
