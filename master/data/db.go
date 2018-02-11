package data

import (
	"gopkg.in/mgo.v2"
)

// NewDb return a new instance of the data layer
func NewDb() (*Db, error) {

	session, err := mgo.Dial("localhost") // TODO: Make this configurable

	if err != nil {
		return nil, err
	}

	return &Db{
		session:      session,
		databaseName: "eaglesight-dev", // TODO: Fetch in config file
	}, nil
}

// Db contains the db
type Db struct {
	session      *mgo.Session
	databaseName string
}

func (db *Db) newDB() *mgo.Database {
	return db.session.Copy().DB(db.databaseName)
}
