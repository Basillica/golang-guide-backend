package userroles

import "gorm.io/gorm"

var (
	ROLES = []string{"GLOBAL_ADMIN", "SUPER_ADMIN", "REGULAR_USER"}
)

type SuperAdmin struct {
	Role       string   // GLOBAL_ADMIN
	UserScopes []string // scopes that define a user's permission
	Roles      []string // list of possible roles in the application
}

// org1
func NewSuperAdmin(r Role, scopes []string) *SuperAdmin {
	return &SuperAdmin{
		Role:       r.String(),
		Roles:      ROLES,
		UserScopes: scopes,
	}
}

func (ga *SuperAdmin) UserPermissionRead() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func (ga *SuperAdmin) UserPermissionWrite() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func (ga *SuperAdmin) ParseUserScopes(db *gorm.DB, baseScopes ...func(db *gorm.DB) *gorm.DB) []func(db *gorm.DB) *gorm.DB {
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, baseScopes...)
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db
	})
	return scopes
}
