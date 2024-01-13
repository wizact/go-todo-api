package db

type SqliteRepository[T any] interface {
	*T
	Connection(*SqliteConnection)
	GetConnection() *SqliteConnection
}

type SqliteRepositoryFactory[T any, P SqliteRepository[T]] struct {
	connection *SqliteConnection
}

func (r *SqliteRepositoryFactory[T, P]) Get() (P, error) {
	if r.connection == nil {
		sc, err := NewSqliteConnection("")

		if err != nil {
			return nil, err
		}
		r.connection = sc
	}

	var result P = new(T)

	result.Connection(r.connection)

	return result, nil
}
