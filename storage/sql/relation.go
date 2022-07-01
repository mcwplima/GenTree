package sql

import (
	"database/sql"
	"fmt"

	model "gentree/models"

	//Postgres Driver
	_ "github.com/lib/pq"
	"github.com/rs/xid"
)

//RelationStore is database persistence interface
type RelationStore struct {
	DB    *sql.DB
	Table string
}

//RelationGet model
func (db RelationStore) RelationGet(objectid string) (*model.Relation, error) {
	var res model.Relation
	row := db.DB.QueryRow(fmt.Sprintf("SELECT objectid, parent_objectid, child_objectid FROM %s WHERE objectid=$1", db.Table), objectid)
	err := row.Scan(&res.ObjectID, &res.Parent, &res.Child)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &res, nil
}

//RelationAdd model
func (db RelationStore) RelationAdd(data *model.Relation) (*model.ObjectID, error) {
	var query string
	resp := model.ObjectID{}
	if data.ObjectID == "" {
		objectid := xid.New()
		data.ObjectID = objectid.String()
		query = fmt.Sprintf("INSERT INTO %s (objectid, parent_objectid, child_objectid) VALUES ($1,$2,$3)", db.Table)
		rows, err := db.DB.Query(query, data.ObjectID, data.Parent, data.Child)
		if err != nil {
			return nil, err
		}
		rows.Close()
	} else {
		query = fmt.Sprintf("UPDATE %s SET parent_objectid=$1 child_objectid=$2 WHERE objectid=$3", db.Table)
		rows, err := db.DB.Query(query, data.Parent, data.Child, data.ObjectID)
		if err != nil {
			return nil, err
		}
		rows.Close()
	}
	resp.ObjectID = data.ObjectID
	return &resp, nil
}

//RelationDel model
func (db RelationStore) RelationDel(objectid string) error {
	rows, err := db.DB.Query(fmt.Sprintf("DELETE FROM %s WHERE objectid=$1", db.Table), objectid)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

//RelationList model
func (db RelationStore) RelationList() ([]*model.Relation, error) {
	results := []*model.Relation{}
	rows, err := db.DB.Query(fmt.Sprintf("SELECT objectid, parent_objectid, child_objectid FROM %s", db.Table))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := model.Relation{}
		err = rows.Scan(&res.ObjectID, &res.Parent, &res.Child)
		if err != nil {
			rows.Close()
			return nil, err
		}
		results = append(results, &res)
	}
	rows.Close()
	return results, nil
}

//RelationList model
func (db RelationStore) RelationListByParent(objectid string) ([]*model.Relation, error) {
	results := []*model.Relation{}
	rows, err := db.DB.Query(fmt.Sprintf("SELECT objectid, parent_objectid, child_objectid FROM %s WHERE parent_objectid = $1", db.Table), objectid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := model.Relation{}
		err = rows.Scan(&res.ObjectID, &res.Parent, &res.Child)
		if err != nil {
			rows.Close()
			return nil, err
		}
		results = append(results, &res)
	}
	rows.Close()
	return results, nil
}

//RelationList model
func (db RelationStore) RelationListByChild(objectid string) ([]*model.Relation, error) {
	results := []*model.Relation{}
	rows, err := db.DB.Query(fmt.Sprintf("SELECT objectid, parent_objectid, child_objectid FROM %s WHERE child_objectid = $1", db.Table))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := model.Relation{}
		err = rows.Scan(&res.ObjectID, &res.Parent, &res.Child)
		if err != nil {
			rows.Close()
			return nil, err
		}
		results = append(results, &res)
	}
	rows.Close()
	return results, nil
}

//RelationCount return the number of services
func (db RelationStore) RelationCount() (int, error) {
	var count int
	rows, err := db.DB.Query(fmt.Sprintf("SELECT COUNT(*) as count FROM %s", db.Table))
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	rows.Close()
	return count, nil
}
