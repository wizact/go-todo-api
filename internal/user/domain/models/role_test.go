package model

import (
	"reflect"
	"testing"
)

func TestRole_IsValid(t *testing.T) {
	type ft func() Role
	r := RoleFactory{}
	tests := []struct {
		name string
		fn   ft
		want bool
	}{
		{
			"Undefined Role",
			r.NewUndefinedRole,
			false,
		},
		{
			"Limited Role",
			r.NewLimitedRole,
			true,
		},
		{
			"Standard Role",
			r.NewStandardRole,
			true,
		},
		{
			"Admin Role",
			r.NewAdminRole,
			true,
		},
		{
			"SuperAdmin Role",
			r.NewSuperAdminRole,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.fn()
			if got := r.IsValid(); got != tt.want {
				t.Errorf("Role.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleType_String(t *testing.T) {
	tests := []struct {
		name string
		r    RoleType
		want string
	}{
		{"NotDefined.String()", NotDefined, "NotDefined"},
		{"Limited.String()", Limited, "Limited"},
		{"Standard.String()", Standard, "Standard"},
		{"Admin.String()", Admin, "Admin"},
		{"SuperAdmin.String()", SuperAdmin, "SuperAdmin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("RoleType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleFactory_NewUndefinedRole(t *testing.T) {
	type ft func() Role
	r := RoleFactory{}
	tests := []struct {
		name string
		fn   ft
		want RoleType
	}{
		{
			"NewUndefinedRole()",
			r.NewUndefinedRole,
			NotDefined,
		},
		{
			"NewLimitedRole()",
			r.NewLimitedRole,
			Limited,
		},
		{
			"Standard()",
			r.NewStandardRole,
			Standard,
		},
		{
			"Admin()",
			r.NewAdminRole,
			Admin,
		},
		{
			"SuperAdmin()",
			r.NewSuperAdminRole,
			SuperAdmin,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(); !reflect.DeepEqual(got.Name, tt.want) {
				t.Errorf("RoleFactory.%v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
