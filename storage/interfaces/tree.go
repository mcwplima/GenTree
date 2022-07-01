package interfaces

import (
	model "gentree/models"
)

//Tree interface
type Tree interface {
	TreeGet(string) ([]*model.PersonTree, error)
}
