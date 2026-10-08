package db

type SQLiteRepository[T any] interface {
	*T
	SetConnection(*SQLiteConnection)
	Connection() *SQLiteConnection
}

type SQLiteRepositoryFactory[T any, P SQLiteRepository[T]] struct {
	connection *SQLiteConnection
}

func (f *SQLiteRepositoryFactory[T, P]) Create() (P, error) {
	if f.connection == nil {
		sc, err := NewSQLiteConnection("")

		if err != nil {
			return nil, err
		}
		f.connection = sc
	}

	var result P = new(T)

	result.SetConnection(f.connection)

	return result, nil
}
