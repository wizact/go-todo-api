package db

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqliteConnectionErrorState struct {
	IsConnectionFailure bool
	HasDriverCause      bool
}

func TestSQLiteSetConnection_Open_PreservesDriverCause(t *testing.T) {
	t.Parallel()

	connection, err := NewSQLiteConnection(filepath.Join(t.TempDir(), "missing", "users.db"))
	if err != nil {
		t.Fatalf("NewSqliteConnection() error = %v", err)
	}

	_, err = connection.Open(gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	var driverError sqlite3.Error
	got := sqliteConnectionErrorState{
		IsConnectionFailure: errors.Is(err, ErrConnectDB),
		HasDriverCause:      errors.As(err, &driverError),
	}
	want := sqliteConnectionErrorState{IsConnectionFailure: true, HasDriverCause: true}

	if got != want {
		t.Fatalf("Open() error state = %#v, want %#v; error = %v", got, want, err)
	}
}
