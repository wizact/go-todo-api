package models

import (
	"time"

	"github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	"github.com/wizact/go-todo-api/internal/user/domain/model"
)

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	Email       string `json:"email"`

	PhoneCountryCode string `json:"phone_country_code"`
	PhoneAreaCode    string `json:"phone_area_code"`
	PhoneNumber      string `json:"phone_number"`
}

func (u *User) ToDomainModel() (aggregate.User, error) {
	var ua aggregate.User = aggregate.NewUser()

	var um model.User = model.NewUser()

	um.FirstName = u.FirstName
	um.LastName = u.FirstName

	if t, e := time.Parse(time.RFC3339, u.DateOfBirth); e != nil {
		um.DateOfBirth = t
	}

	um.Email = u.Email

	um.Phone.CountryCode = u.PhoneCountryCode
	um.Phone.AreaCode = u.PhoneAreaCode
	um.Phone.Number = u.PhoneNumber

	return ua, nil
}

func (u *User) ToApiModel(ua aggregate.User) error {
	u.FirstName = ua.User().FirstName
	u.LastName = ua.User().FirstName
	u.DateOfBirth = ua.User().DateOfBirth.String()
	u.Email = ua.User().Email

	u.PhoneCountryCode = ua.User().Phone.CountryCode
	u.PhoneAreaCode = ua.User().Phone.AreaCode
	u.PhoneNumber = ua.User().Phone.Number

	return nil
}
