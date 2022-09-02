package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
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

func (a *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello")
	}
}
