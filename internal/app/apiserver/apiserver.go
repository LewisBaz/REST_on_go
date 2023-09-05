package apiserver

import (
	"database/sql"
	"main/internal/app/store/sqlstore"
	"net/http"
)

func Start(c *Config) error {

	db, err := newDB(c.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(c.BindAddress, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}