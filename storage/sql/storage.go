package sql

import (
	//Postgres Driver
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"gentree/config"
	"gentree/storage/interfaces"
)

var (
	personTable   = "persons"
	relationTable = "relations"
)

//Store implements the storage interface
type Store struct {
	person   *PersonStore
	relation *RelationStore
	tree     *TreeStore

	db *sql.DB
}

//Person returns the memory implementation for interfaces.Person
func (s *Store) Person() interfaces.Person {
	if s.person == nil {
		s.person = &PersonStore{DB: s.db, Table: personTable}
	}
	return s.person
}

func (s *Store) Relation() interfaces.Relation {
	if s.relation == nil {
		s.relation = &RelationStore{DB: s.db, Table: relationTable}
	}
	return s.relation
}

func (s *Store) Tree() interfaces.Tree {
	if s.tree == nil {
		s.tree = &TreeStore{DB: s.db, Table: relationTable}
	}
	return s.tree
}

// FromConfig instantiates for a given configuration
func FromConfig(c *config.Config) *Store {
	server := c.Database.Server
	port := c.Database.Port
	dbname := c.Database.Database
	user := c.Database.User
	password := c.Database.Password

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, dbname)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Store{db: db}
}
