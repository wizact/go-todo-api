package model

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestRehydrateUser_PreservesID(t *testing.T) {
	t.Parallel()

	want := uuid.New()
	user := RehydrateUser(
		want,
		"Ada",
		"Lovelace",
		time.Date(1815, time.December, 10, 0, 0, 0, 0, time.UTC),
		"ada@example.com",
		NewPhoneNumber("+44", "20", "12345678"),
	)

	if got := user.ID(); got != want {
		t.Fatalf("user ID = %v, want %v", got, want)
	}
}

func TestNewUser_InvalidDetails_ReturnsError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		firstName string
		lastName  string
		email     string
	}{
		{name: "missing name", email: "ada@example.com"},
		{name: "invalid email", firstName: "Ada", lastName: "Lovelace", email: "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewUser(tt.firstName, tt.lastName, time.Time{}, tt.email, PhoneNumber{})
			if !errors.Is(err, ErrInvalidUser) {
				t.Fatalf("error = %v, want %v", err, ErrInvalidUser)
			}
		})
	}
}

func TestNewUser_ValidDetails_GeneratesID(t *testing.T) {
	t.Parallel()

	user, err := NewUser("Ada", "Lovelace", time.Time{}, "ada@example.com", PhoneNumber{})
	if err != nil {
		t.Fatalf("NewUser() error = %v", err)
	}

	if got := user.ID(); got == uuid.Nil {
		t.Fatal("user ID is empty")
	}
}

func TestUser_DoesNotExposeConstructionMutation(t *testing.T) {
	t.Parallel()

	typeOfUser := reflect.TypeFor[*User]()
	for _, methodName := range []string{"SetID", "SetName", "SetDateOfBirth", "SetEmail", "SetPhone"} {
		t.Run(methodName, func(t *testing.T) {
			t.Parallel()

			_, exposed := typeOfUser.MethodByName(methodName)
			if exposed {
				t.Fatalf("User exposes %s", methodName)
			}
		})
	}
}

func TestUser_IsValid(t *testing.T) {
	type user struct {
		ID          uuid.UUID
		FirstName   string
		LastName    string
		DateOfBirth time.Time
		Email       string
		Phone       PhoneNumber
	}
	tests := []struct {
		name   string
		fields user
		want   bool
	}{
		{"invalid user with no email", user{FirstName: "foo", LastName: "bar", Email: ""}, false},
		{"invalid user with no valid email", user{FirstName: "foo", LastName: "bar", Email: "invalidemail"}, false},
		{"invalid user with no name", user{FirstName: "", LastName: "", Email: "foo@bar.baz"}, false},
		{"valid user", user{FirstName: "foo", LastName: "bar", Email: "foo@bar.baz"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := RehydrateUser(
				tt.fields.ID,
				tt.fields.FirstName,
				tt.fields.LastName,
				tt.fields.DateOfBirth,
				tt.fields.Email,
				tt.fields.Phone,
			)
			if got := u.IsValid(); got != tt.want {
				t.Errorf("User.IsValid() = %v, want %v %v", got, tt.want, u)
			}
		})
	}
}

func TestUser_IsTheSameUserAs(t *testing.T) {
	type user struct {
		ID          uuid.UUID
		FirstName   string
		LastName    string
		DateOfBirth time.Time
		Email       string
		Phone       PhoneNumber
	}

	user2 := RehydrateUser(uuid.New(), "foo", "bar", time.Time{}, "foo@bar.baz", PhoneNumber{})

	tests := []struct {
		name   string
		fields user
		want   bool
	}{
		{"Users with the same id", user{ID: user2.ID(), FirstName: "foo", LastName: "bar", Email: "foo@bar.baz"}, true},
		{"Users with the different id", user{ID: uuid.New(), FirstName: "foo", LastName: "bar", Email: "foo@bar.baz"}, false},
		{"Users with the different id and same email", user{ID: uuid.New(), FirstName: "foo", LastName: "bar", Email: "foo@bar.baz"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := RehydrateUser(
				tt.fields.ID,
				tt.fields.FirstName,
				tt.fields.LastName,
				tt.fields.DateOfBirth,
				tt.fields.Email,
				tt.fields.Phone,
			)
			if got := u.IsTheSameUserAs(user2); got != tt.want {
				t.Errorf("User.IsTheSameUserAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasName(t *testing.T) {
	user1 := User{}
	user2 := RehydrateUser(uuid.Nil, "foo", "", time.Time{}, "", PhoneNumber{})
	user3 := RehydrateUser(uuid.Nil, "", "bar", time.Time{}, "", PhoneNumber{})
	tests := []struct {
		name string
		user User
		want bool
	}{
		{"user with no FirstName and LastName", user1, false},
		{"user with FirstName but no LastName", user2, true},
		{"user with LastName but no FirstName", user3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasName(tt.user); got != tt.want {
				t.Errorf("HasName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"user with no email", "", false},
		{"user with no valid email", "invalidemail", false},
		{"user with valid email", "foo@bar.baz", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := RehydrateUser(uuid.Nil, "", "", time.Time{}, tt.email, PhoneNumber{})
			if got := HasValidEmail(u); got != tt.want {
				t.Errorf("HasValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_IsEqual(t *testing.T) {
	type fields struct {
		CountryCode string
		AreaCode    string
		Number      string
	}
	p2 := NewPhoneNumber("+64", "021", "123456")

	tests := []struct {
		name   string
		fields fields
		phone  PhoneNumber
		want   bool
	}{
		{"valid phone number", fields{CountryCode: "+64", AreaCode: "021", Number: "123456"}, p2, true},
		{"country code mismatch", fields{CountryCode: "+61", AreaCode: "021", Number: "123456"}, p2, false},
		{"area code mismatch", fields{CountryCode: "+64", AreaCode: "022", Number: "123456"}, p2, false},
		{"number mismatch", fields{CountryCode: "+64", AreaCode: "021", Number: "123457"}, p2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPhoneNumber(
				tt.fields.CountryCode,
				tt.fields.AreaCode,
				tt.fields.Number,
			)
			if got := p.IsEqual(tt.phone); got != tt.want {
				t.Errorf("PhoneNumber.IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_DoesNotExposeMutation(t *testing.T) {
	t.Parallel()

	typeOfPhoneNumber := reflect.TypeFor[*PhoneNumber]()
	for _, methodName := range []string{"SetCountryCode", "SetAreaCode", "SetNumber"} {
		t.Run(methodName, func(t *testing.T) {
			t.Parallel()

			_, exposed := typeOfPhoneNumber.MethodByName(methodName)
			if exposed {
				t.Fatalf("PhoneNumber exposes %s", methodName)
			}
		})
	}
}

func TestNewLocation_InitializesCoordinates(t *testing.T) {
	t.Parallel()

	want := [2]float64{173.3002574488138, -41.26595602617756}
	location := NewLocation(want[0], want[1])
	longitude, latitude := location.Coordinates()

	if got := [2]float64{longitude, latitude}; got != want {
		t.Fatalf("coordinates = %v, want %v", got, want)
	}
}

func TestLocation_DoesNotExposeCoordinateMutation(t *testing.T) {
	t.Parallel()

	_, exposed := reflect.TypeFor[*Location]().MethodByName("SetCoordinates")
	if exposed {
		t.Fatal("Location exposes SetCoordinates")
	}
}

func TestLocation_DoesNotExposeFields(t *testing.T) {
	t.Parallel()

	typeOfLocation := reflect.TypeFor[Location]()
	for index := range typeOfLocation.NumField() {
		if field := typeOfLocation.Field(index); field.IsExported() {
			t.Fatalf("Location exposes field %s", field.Name)
		}
	}
}
