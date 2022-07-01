package sql

import (
	"database/sql"
	model "gentree/models"
	"gentree/util"
)

type TreeStore struct {
	DB    *sql.DB
	Table string
}

//RelationList model
func (db TreeStore) TreeGet(objectid string) ([]*model.PersonTree, error) {
	results := []*model.PersonTree{}
	rows, err := db.DB.Query(util.TreeQueryBuilder(objectid))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := model.PersonTree{}
		err = rows.Scan(&res.ObjectID, &res.Name)
		if err != nil {
			rows.Close()
			return nil, err
		}
		results = append(results, &res)
	}
	rows.Close()
	return results, nil
}
