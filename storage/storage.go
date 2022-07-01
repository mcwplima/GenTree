package storage

import (
	"context"

	"gentree/config"
	"gentree/storage/interfaces"
	"gentree/storage/sql"
)

type key int

const storeKey key = 99

type Store interface {
	Person() interfaces.Person
	Relation() interfaces.Relation
	Tree() interfaces.Tree
}

// FromConfig returns an storage for the given DB_TYPE
func FromConfig(c *config.Config) Store {
	return sql.FromConfig(c)
}

// NewContext generates a new Context storing the Store into its values.
// Thats helpfull if you need to transfer the store inside the context
// to another function.
func NewContext(ctx context.Context, s Store) context.Context {
	return context.WithValue(ctx, storeKey, s)
}

// FromContext retrieves a *Store previously added to the context by the NewContext func.
// It returns nil if no Store is found.
func FromContext(ctx context.Context) (Store, bool) {
	s, ok := ctx.Value(storeKey).(Store)
	return s, ok
}
