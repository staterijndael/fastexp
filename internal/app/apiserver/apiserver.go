package apiserver

import (
	"net/http"

	"github.com/Oringik/fastexp/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	srv.logger.Info("Server starting")

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
