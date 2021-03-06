package db

import (
	"errors"
	"time"

	"github.com/globalsign/mgo"
)

type Database struct {
	db *mgo.Database
}

func (d *Database) Close() {
	d.db.Session.Close()
}

func (d *Database) C(name string) *mgo.Collection {
	return d.db.C(name)
}

func (d *Database) Run(cmd, result interface{}) error {
	return d.db.Run(cmd, result)
}

type Dao struct {
	session *mgo.Session
}

func New(addr string) (*Dao, error) {
	if addr == "" {
		return nil, errors.New("database address must be configured for mongo storage")
	}

	s, err := mgo.DialWithTimeout(addr, time.Second*5)
	if err != nil {
		return nil, err
	}

	d := &Dao{
		session: s,
	}
	return d, nil
}

func (d *Dao) Close() {
	d.session.Close()
}

func (d *Dao) db() *Database {
	return &Database{
		db: d.session.Copy().DB(""),
	}
}

func (d *Dao) Do(fn func(db *Database)) {
	db := d.db()
	defer db.Close()

	fn(db)
}
