package models

import (
	"net/http"
	"time"

	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	model "github.com/wizact/go-todo-api/internal/user/domain/models"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

type User struct {
	UserID      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	Email       string `json:"email"`

	PhoneCountryCode string `json:"phone_country_code"`
	PhoneAreaCode    string `json:"phone_area_code"`
	PhoneNumber      string `json:"phone_number"`

	LocationLongitude float64 `json:"location_longitude"`
	LocationLatitude  float64 `json:"location_latitude"`
}

func (u *User) ToDomainModel() (aggregate.User, *hsm.AppError) {
	ua := aggregate.NewUser()

	dateOfBirth, err := time.Parse(time.RFC3339, u.DateOfBirth)
	if err != nil {
		return ua, &hsm.AppError{SanitisedMessage: err.Error(), ErrorObject: err, Code: http.StatusBadRequest}
	}

	phone := model.NewPhoneNumber(u.PhoneCountryCode, u.PhoneAreaCode, u.PhoneNumber)
	duser, err := model.NewUser(u.FirstName, u.LastName, dateOfBirth, u.Email, phone)
	if err != nil {
		return ua, hsm.NewAppError(err, "user info is not valid", http.StatusBadRequest)
	}

	ua.SetUser(duser)

	dloc := model.NewLocation(u.LocationLongitude, u.LocationLatitude)
	ua.SetLocation(dloc)

	return ua, nil
}

func (u *User) ToApiModel(ua aggregate.User) *hsm.AppError {
	uau := ua.User()
	fn, ln := uau.Name()
	u.UserID = uau.ID().String()
	u.FirstName = fn
	u.LastName = ln
	u.DateOfBirth = uau.DateOfBirth().String()
	u.Email = uau.Email()

	uaup := uau.Phone()
	u.PhoneCountryCode = uaup.CountryCode()
	u.PhoneAreaCode = uaup.AreaCode()
	u.PhoneNumber = uaup.Number()

	location := ua.Location()
	u.LocationLongitude, u.LocationLatitude = location.Coordinates()

	return nil
}
