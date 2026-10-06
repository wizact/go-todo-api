package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wizact/go-todo-api/internal/user/application/registration"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	repositoryport "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SqliteRegistrationVerification struct {
	UserID       string `gorm:"primaryKey;not null"`
	SecretDigest string `gorm:"not null"`
	ExpiresAt    int64  `gorm:"not null"`
	CreatedAt    int64  `gorm:"autoCreateTime:milli;not null"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli;not null"`
}

func (r *UserSqliteRepository) CompleteRegistration(
	ctx context.Context,
	user aggregate.User,
	verificationDigest string,
) (aggregate.User, error) {
	emptyUser := aggregate.User{}
	database, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return emptyUser, fmt.Errorf("open registration completion database: %w", err)
	}

	err = database.WithContext(ctx).Transaction(func(transaction *gorm.DB) error {
		result := transaction.
			Where("user_id = ? AND secret_digest = ?", user.UserId().String(), verificationDigest).
			Delete(&SqliteRegistrationVerification{})
		if result.Error != nil {
			return fmt.Errorf("consume registration verification: %w", result.Error)
		}
		if result.RowsAffected != 1 {
			return repositoryport.ErrRegistrationVerificationNotFound
		}

		persistedAggregate := SqliteUserAggregate{UserID: user.UserId().String()}
		if err := transaction.First(&persistedAggregate).Error; err != nil {
			return fmt.Errorf("find user aggregate for registration completion: %w", err)
		}
		updatedAggregate := SqliteUserAggregate{}
		updatedAggregate.FromDomainEntityToDbModel(user)
		persistedAggregate.ValueData = updatedAggregate.ValueData
		if err := transaction.Save(&persistedAggregate).Error; err != nil {
			return fmt.Errorf("update user aggregate for registration completion: %w", err)
		}

		emailView := SqliteUserEmailView{UserID: user.UserId().String()}
		if err := transaction.First(&emailView).Error; err != nil {
			return fmt.Errorf("find user email view for registration completion: %w", err)
		}
		emailView.HasVerifiedEmail = user.HasVerifiedEmail()
		if err := transaction.Save(&emailView).Error; err != nil {
			return fmt.Errorf("update user email view for registration completion: %w", err)
		}

		return nil
	})
	if err != nil {
		return emptyUser, fmt.Errorf("complete user registration: %w", err)
	}

	return user, nil
}

// TableName overrides GORM's default table name.
func (SqliteRegistrationVerification) TableName() string {
	return "user_registration_verifications"
}

func (record SqliteRegistrationVerification) toDomain() (registration.Verification, error) {
	userID, err := uuid.Parse(record.UserID)
	if err != nil {
		return registration.Verification{}, fmt.Errorf("parse persisted registration verification user ID: %w", err)
	}

	return registration.Verification{
		UserID:       userID,
		SecretDigest: record.SecretDigest,
		ExpiresAt:    time.UnixMilli(record.ExpiresAt).UTC(),
	}, nil
}

func (r *UserSqliteRepository) SaveRegistrationVerification(
	ctx context.Context,
	verification registration.Verification,
) error {
	database, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return fmt.Errorf("open registration verification database: %w", err)
	}

	record := SqliteRegistrationVerification{
		UserID:       verification.UserID.String(),
		SecretDigest: verification.SecretDigest,
		ExpiresAt:    verification.ExpiresAt.UnixMilli(),
	}
	result := database.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"secret_digest", "expires_at", "updated_at"}),
		}).
		Create(&record)
	if result.Error != nil {
		return fmt.Errorf("save registration verification: %w", result.Error)
	}

	return nil
}

func (r *UserSqliteRepository) FindRegistrationVerification(
	ctx context.Context,
	userID uuid.UUID,
) (registration.Verification, error) {
	emptyVerification := registration.Verification{}
	database, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return emptyVerification, fmt.Errorf("open registration verification database: %w", err)
	}

	record := SqliteRegistrationVerification{UserID: userID.String()}
	if err := database.WithContext(ctx).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return emptyVerification, repositoryport.ErrRegistrationVerificationNotFound
		}
		return emptyVerification, fmt.Errorf("find registration verification: %w", err)
	}

	verification, err := record.toDomain()
	if err != nil {
		return emptyVerification, fmt.Errorf("map registration verification: %w", err)
	}

	return verification, nil
}
