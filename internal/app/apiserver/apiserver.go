package apiserver

import (
	"TrainingRESTAPI/internal/app/store"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	if err := s.configurteStore(); err != nil{
		return err
	}

	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BinAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configurteStore() error {
	s.config.Store.DatabaseURL = fmt.Sprintf(s.config.Store.DatabaseURL , os.Getenv("PWDDB"))
	print(s.config.Store.DatabaseURL)
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil{
		return err
	}
	s.store = st
	return nil
}


func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())

}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}


