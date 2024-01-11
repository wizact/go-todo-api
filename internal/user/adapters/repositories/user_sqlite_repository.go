package repository

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserSqliteRepository struct {
}

func NewUserSqlLiteRepository() *UserSqliteRepository {
	return &UserSqliteRepository{}
}

func (r *UserSqliteRepository) FindById(ctx context.Context, id uuid.UUID) (ua.User, error) {
	return ua.User{}, nil
}

func (r *UserSqliteRepository) FindByEmail(ctx context.Context, email string) (ua.User, error) {
	return ua.User{}, nil
}

func (r *UserSqliteRepository) Create(ctx context.Context, user ua.User) (ua.User, error) {
	u := &SqliteUserAggregate{}
	u.FromDomainEntityToDbModel(user)

	db, err := gorm.Open(sqlite.Open("/go/src/go-todo-api/db/todo.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	result := db.Create(&u)

	log.Println(result.Error, result.RowsAffected)

	return ua.User{}, nil
}

func (r *UserSqliteRepository) Update(ctx context.Context, user ua.User) (ua.User, error) {
	return ua.User{}, nil
}

type SqliteUserAggregate struct {
	UserID    string          `gorm:"primaryKey;not null"`
	ValueData SqliteUserModel `gorm:"serializer:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
}

func (dbm *SqliteUserAggregate) FromDomainEntityToDbModel(de ua.User) {
	log.Println(de)
	dbm.UserID = de.UserId().String()
	dbm.ValueData = SqliteUserModel{
		ID:          de.UserId().String(),
		FirstName:   de.User().FirstName,
		LastName:    de.User().LastName,
		DateOfBirth: de.User().DateOfBirth,
		Email:       de.User().Email,
		CountryCode: de.User().Phone.CountryCode,
		AreaCode:    de.User().Phone.AreaCode,
		Number:      de.User().Phone.Number,
	}
}

func (dbm SqliteUserAggregate) FromDbModelToDomainEntity() ua.User {
	log.Println(dbm)
	return ua.User{}
}

// TableName overrides grom default table name
func (SqliteUserAggregate) TableName() string {
	return "users_aggregate"
}
