package userroles

import "github.com/Basillica/golang-guide/service"

type Role int64

const (
	SUPER_ADMIN Role = iota
	ORG_ADMIN
	REGULAR_USER
	UNIMPLEMENTED_ROLE
)

func (r Role) String() string {
	switch r {
	case SUPER_ADMIN:
		return "SUPER_ADMIN"
	case ORG_ADMIN:
		return "ORG_ADMIN"
	case REGULAR_USER:
		return "REGULAR_USER"
	}
	return "UNIMPLEMENTED_ROLE"
}

func (Role) Role(role string) Role {
	switch role {
	case "SUPER_ADMIN":
		return SUPER_ADMIN
	case "ORG_ADMIN":
		return ORG_ADMIN
	case "REGULAR_USER":
		return REGULAR_USER
	}
	return UNIMPLEMENTED_ROLE
}

func (r Role) RoleModel(scopes []string) (service.UserRoleModel, error) {
	switch r {
	case SUPER_ADMIN:
		return NewSuperAdmin(r, scopes), nil
	case ORG_ADMIN:
		return NewOrganisationAdmin(r, scopes), nil
	case REGULAR_USER:
		return NewRegularUser(r, scopes), nil
	}
	return nil, nil
}
