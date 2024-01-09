package models

import (
	"net/http"
	"time"

	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	model "github.com/wizact/go-todo-api/internal/user/domain/models"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
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

func (u *User) ToDomainModel() (aggregate.User, *hsm.AppError) {
	var ua aggregate.User = aggregate.NewUser()

	duser := model.NewUser()

	duser.FirstName = u.FirstName
	duser.LastName = u.FirstName

	if t, e := time.Parse(time.RFC3339, u.DateOfBirth); e != nil {
		duser.DateOfBirth = t
	} else {
		return ua, &hsm.AppError{Message: e.Error(), ErrorObject: e, Code: http.StatusBadRequest}
	}

	duser.Email = u.Email

	duser.Phone.CountryCode = u.PhoneCountryCode
	duser.Phone.AreaCode = u.PhoneAreaCode
	duser.Phone.Number = u.PhoneNumber

	ua.SetUser(duser)

	return ua, nil
}

func (u *User) ToApiModel(ua aggregate.User) *hsm.AppError {
	u.FirstName = ua.User().FirstName
	u.LastName = ua.User().FirstName
	u.DateOfBirth = ua.User().DateOfBirth.String()
	u.Email = ua.User().Email

	u.PhoneCountryCode = ua.User().Phone.CountryCode
	u.PhoneAreaCode = ua.User().Phone.AreaCode
	u.PhoneNumber = ua.User().Phone.Number

	return nil
}
