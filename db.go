package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Store interface {
	ListExhibitions() ([]Exhibition, error)
	GetExhibition(id int8) (Exhibition, error)
	CreateExhibition() (Exhibition, error)
	EditExhibitions() (Exhibition, error)
	DeleteExhibition() error
}

func NewStore(db *sql.DB) Store {
	return &DB{conn: db}
}

type DB struct {
	conn *sql.DB
}

func (db *DB) ListExhibitions() ([]Exhibition, error) {
	var exhibitions []Exhibition
	var err error
	return exhibitions, err
}

func (db *DB) GetExhibition(id int8) (exhibition Exhibition, err error) {
	err = db.conn.QueryRow("SELECT * FROM exhibitions WHERE id=? LIMIT 1", id).Scan(
		&exhibition.ID,
		&exhibition.Name,
	)
	return
}

func (db *DB) CreateExhibition() (Exhibition, error) {
	var exhibition Exhibition
	var err error
	return exhibition, err
}

func (db *DB) EditExhibitions() (Exhibition, error) {
	var exhibition Exhibition
	var err error
	return exhibition, err
}

func (db *DB) DeleteExhibition() error {
	var err error
	return err
}
