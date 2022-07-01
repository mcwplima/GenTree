package interfaces

import (
	model "gentree/models"
)

//Relation interface
type Relation interface {
	RelationGet(string) (*model.Relation, error)
	RelationAdd(*model.Relation) (*model.ObjectID, error)
	RelationDel(string) error
	RelationList() ([]*model.Relation, error)
	RelationListByParent(string) ([]*model.Relation, error)
	RelationListByChild(string) ([]*model.Relation, error)
	RelationCount() (int, error)
}
