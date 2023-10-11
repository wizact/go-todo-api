package model

import (
	"github.com/google/uuid"
	sp "github.com/wizact/go-todo-api/pkg/specification"
)

type RoleType int

const (
	NotDefined RoleType = iota
	Limited
	Standard
	Admin
	SuperAdmin
)

type Role struct {
	ID   uuid.UUID
	Name RoleType
}

func NewRole(r RoleType) Role {
	return Role{ID: uuid.New(), Name: r}
}

func (r Role) IsValid() bool {
	spec := sp.NewAndSpecification[Role](
		sp.FunctionSpecification[Role](func(u Role) bool { return u.Name != NotDefined }),
	)

	return spec.IsValid(r)

}

func (r RoleType) String() string {
	return [...]string{"NotDefined", "Limited", "Standard", "Admin", "SuperAdmin"}[r]
}

type RoleFactory struct{}

func (r RoleFactory) NewUndefinedRole() Role {
	return NewRole(Limited)
}

func (r RoleFactory) NewLimitedRole() Role {
	return NewRole(Limited)
}

func (r RoleFactory) NewStandardRole() Role {
	return NewRole(Standard)
}

func (r RoleFactory) NewAdminRole() Role {
	return NewRole(Limited)
}

func (r RoleFactory) NewSuperAdminRole() Role {
	return NewRole(Limited)
}
