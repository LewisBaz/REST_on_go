package apiserver

import (
	"database/sql"
	"main/internal/app/store/sqlstore"
	"net/http"

	"github.com/gorilla/sessions"
)

func Start(c *Config) error {

	db, err := newDB(c.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionsStore := sessions.NewCookieStore([]byte(c.SessionKey))
	srv := newServer(store, sessionsStore)

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