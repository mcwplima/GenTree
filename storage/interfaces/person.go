package interfaces

import (
	model "gentree/models"
)

//Person interface
type Person interface {
	PersonGet(string) (*model.Person, error)
	PersonAdd(*model.Person) (*model.ObjectID, error)
	PersonDel(string) error
	PersonList() ([]*model.Person, error)
	PersonListByChild(string) ([]*model.Person, error)
	PersonCount() (int, error)
}
