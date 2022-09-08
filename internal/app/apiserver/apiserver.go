package apiserver

import (
	"Rest/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *APIServer) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}

	a.configureRouter()

	if err := a.configureStore(); err != nil {
		return err
	}

	a.logger.Info("starting api server")
	return http.ListenAndServe(a.config.BindAddr, a.router)
}

func (a *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.Loglevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(level)

	return nil
}

func (a *APIServer) configureRouter() {
	a.router.HandleFunc("/hello", a.handleHello())
}

func (a *APIServer) configureStore() error {
	st := store.NewStore(a.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	a.store = st

	return nil
}

func (a *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello")
	}
}
