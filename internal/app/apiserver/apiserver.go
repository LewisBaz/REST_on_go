package apiserver

import (
	"io"
	"main/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *APIserver) Start() error {
	if err := api.configureLogger(); err != nil {
		return err
	}

	api.logger.Info("starting api server")

	api.configureRouter()

	if err := api.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.BindAddress, api.router)
}

func (api *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(api.config.LogLevel)
	if err != nil {
		return err
	}

	api.logger.SetLevel(level)

	return nil
}

func (api *APIserver) configureRouter() {
	api.router.HandleFunc("/hello", api.handleHello())
}

func (api *APIserver) configureStore() error {
	str := store.New(api.config.Store)
	if err := str.Open(); err != nil {
		return err
	}

	api.store = str

	return nil
}

func (api *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
