package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	model "github.com/wizact/go-todo-api/internal/user/domain/models"
	us "github.com/wizact/go-todo-api/internal/user/domain/services"
	"gorm.io/gorm"
)

type UserSqliteRepository struct {
	connection *dbinfra.SqliteConnection
}

func (r *UserSqliteRepository) Connection(cnn *dbinfra.SqliteConnection) {
	r.connection = cnn
}

func (r *UserSqliteRepository) GetConnection() *dbinfra.SqliteConnection {
	return r.connection
}

func (r *UserSqliteRepository) FindById(ctx context.Context, id uuid.UUID) (ua.User, error) {
	emptyUser := ua.User{}
	db, err := r.connection.Open(gorm.Config{})

	if err != nil {
		return emptyUser, err
	}

	u := &SqliteUserAggregate{UserID: id.String()}

	result := db.Limit(1).First(u)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return emptyUser, us.ErrUserIdDoesNotExist
	}

	if result.Error != nil {
		return emptyUser, result.Error
	}

	de := u.FromDbModelToDomainEntity()

	return de, nil
}

func (r *UserSqliteRepository) FindByEmail(ctx context.Context, email string) (ua.User, error) {
	emptyUser := ua.User{}
	db, err := r.connection.Open(gorm.Config{})

	if err != nil {
		return emptyUser, err
	}

	uev := &SqliteUserEmailView{Email: email}
	result := db.Where(uev).First(uev)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return emptyUser, us.ErrUserByEmailDoesNotExist
	}

	if result.Error != nil {
		return emptyUser, result.Error
	}

	de := uev.FromDbModelToDomainEntity()

	u, err := r.FindById(ctx, de.Id())

	if err != nil {
		return emptyUser, err
	}

	return u, nil
}

func (r *UserSqliteRepository) Create(ctx context.Context, user ua.User) (ua.User, error) {
	emptyUser := ua.User{}

	db, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return emptyUser, err
	}

	tx := db.Begin()

	u := &SqliteUserAggregate{}
	u.FromDomainEntityToDbModel(user)

	result := tx.Create(&u)

	if result.Error != nil {
		tx.Rollback()
		return emptyUser, result.Error
	}

	user = u.FromDbModelToDomainEntity()
	_, err = r.createOrUpdateUserEmailView(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return emptyUser, result.Error
	}

	tx.Commit()
	return user, nil
}

func (r *UserSqliteRepository) Update(ctx context.Context, user ua.User) (ua.User, error) {
	return ua.User{}, nil
}

type SqliteUserAggregate struct {
	UserID    string          `gorm:"primaryKey;not null"`
	ValueData SqliteUserModel `gorm:"serializer:json"`
	CreatedAt int64           `gorm:"autoCreateTime:milli"`
	UpdatedAt int64           `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt  `gorm:"index"`
}

type SqliteUserModel struct {
	ID          string
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	Email       string
	CountryCode string
	AreaCode    string
	Number      string

	LocationLong float64
	LocationLat  float64

	HasVerifiedEmail bool
	IsActive         bool
}

// TableName overrides grom default table name
func (SqliteUserAggregate) TableName() string {
	return "users_aggregate"
}

func (dbm *SqliteUserAggregate) FromDomainEntityToDbModel(de ua.User) {
	dbm.UserID = de.UserId().String()
	dbm.ValueData = SqliteUserModel{
		ID:               de.UserId().String(),
		FirstName:        de.User().FirstName,
		LastName:         de.User().LastName,
		DateOfBirth:      de.User().DateOfBirth,
		Email:            de.User().Email,
		CountryCode:      de.User().Phone.CountryCode,
		AreaCode:         de.User().Phone.AreaCode,
		Number:           de.User().Phone.Number,
		LocationLong:     de.Location().Longitude,
		LocationLat:      de.Location().Latitude,
		HasVerifiedEmail: de.HasVerifiedEmail(),
		IsActive:         de.IsActive(),
	}
}

func (dbm SqliteUserAggregate) FromDbModelToDomainEntity() ua.User {
	de := ua.NewUser()
	mu := model.User{
		ID:          uuid.MustParse(dbm.UserID),
		FirstName:   dbm.ValueData.FirstName,
		LastName:    dbm.ValueData.LastName,
		DateOfBirth: dbm.ValueData.DateOfBirth,
		Email:       dbm.ValueData.Email,
		Phone: model.PhoneNumber{
			CountryCode: dbm.ValueData.CountryCode,
			AreaCode:    dbm.ValueData.AreaCode,
			Number:      dbm.ValueData.Number,
		},
	}

	dl := model.NewLocation()
	dl.SetCoordinates(dbm.ValueData.LocationLong, dbm.ValueData.LocationLat)

	de.SetHasVerifiedEmail(dbm.ValueData.HasVerifiedEmail)
	de.SetIsActive(dbm.ValueData.IsActive)

	de.SetUser(mu)
	de.SetLocation(dl)

	return de
}
