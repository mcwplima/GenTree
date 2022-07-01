package sql

import (
	"database/sql"
	"fmt"

	model "gentree/models"

	//Postgres Driver
	_ "github.com/lib/pq"
	"github.com/rs/xid"
)

//PersonStore is database persistence interface
type PersonStore struct {
	DB    *sql.DB
	Table string
}

//PersonGet model
func (db PersonStore) PersonGet(objectid string) (*model.Person, error) {
	var res model.Person
	row := db.DB.QueryRow(fmt.Sprintf("SELECT objectid, name FROM %s WHERE objectid=$1", db.Table), objectid)
	err := row.Scan(&res.ObjectID, &res.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &res, nil
}

//PersonAdd model
func (db PersonStore) PersonAdd(data *model.Person) (*model.ObjectID, error) {
	var query string
	resp := model.ObjectID{}
	if data.ObjectID == "" {
		objectid := xid.New()
		data.ObjectID = objectid.String()
		query = fmt.Sprintf("INSERT INTO %s (objectid, name) VALUES ($1,$2)", db.Table)
		rows, err := db.DB.Query(query, data.ObjectID, data.Name)
		if err != nil {
			return nil, err
		}
		rows.Close()
	} else {
		query = fmt.Sprintf("UPDATE %s SET name=$1 WHERE objectid=$2", db.Table)
		rows, err := db.DB.Query(query, data.Name, data.ObjectID)
		if err != nil {
			return nil, err
		}
		rows.Close()
	}
	resp.ObjectID = data.ObjectID
	return &resp, nil
}

//PersonDel model
func (db PersonStore) PersonDel(objectid string) error {
	rows, err := db.DB.Query(fmt.Sprintf("DELETE FROM %s WHERE objectid=$1", db.Table), objectid)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

//PersonList model
func (db PersonStore) PersonList() ([]*model.Person, error) {
	results := []*model.Person{}
	rows, err := db.DB.Query(fmt.Sprintf("SELECT objectid, name FROM %s ORDER BY name", db.Table))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := model.Person{}
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

//RelationList model
func (db PersonStore) PersonListByChild(objectid string) ([]*model.Person, error) {
	results := []*model.Person{}
	rows, err := db.DB.Query(fmt.Sprintf("SELECT p.objectid, p.name FROM %s as p, relations as r WHERE p.objectid = r.parent_objectid AND r.child_objectid = $1", db.Table), objectid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := model.Person{}
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

//PersonCount return the number of services
func (db PersonStore) PersonCount() (int, error) {
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
