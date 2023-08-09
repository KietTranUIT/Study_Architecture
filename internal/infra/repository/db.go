package repository

import (
	"database/sql"
	"user-service/internal/core/port/repository"
	"user-service/internal/infra/repository/config"
)

type database struct {
	*sql.DB
}

func NewDB(conf config.ConfigDatabase) (repository.Database, error) {
	db, err := newDatabase(conf)

	if err != nil {
		return nil, err
	}

	return &database{
		db,
	}, nil
}

func newDatabase(conf config.ConfigDatabase) (*sql.DB, error) {
	db, err := sql.Open(conf.Driver, conf.URL())

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db database) Close() error {
	return db.DB.Close()
}

func (db database) GetDB() *sql.DB {
	return db.DB
}
