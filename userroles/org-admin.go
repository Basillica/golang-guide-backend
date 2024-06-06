package userroles

import "gorm.io/gorm"

type OrganisationAdmin struct {
	Role       string   // GLOBAL_ADMIN
	UserScopes []string // scopes that define a user's permission
	Roles      []string // list of possible roles in the application
}

// org1
func NewOrganisationAdmin(r Role, scopes []string) *OrganisationAdmin {
	return &OrganisationAdmin{
		Role:       r.String(),
		Roles:      ROLES,
		UserScopes: scopes,
	}
}

func (ga *OrganisationAdmin) UserPermissionRead() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("scope LIKE ?", `%"`+ga.UserScopes[0]+`%`)
	}
}

func (ga *OrganisationAdmin) UserPermissionWrite() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("scope LIKE ?", `%"`+ga.UserScopes[0]+`%`)
	}
}

func (ga *OrganisationAdmin) ParseUserScopes(db *gorm.DB, baseScopes ...func(db *gorm.DB) *gorm.DB) []func(db *gorm.DB) *gorm.DB {
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, baseScopes...)
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db
	})
	return scopes
}
