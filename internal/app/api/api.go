package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/metall27/ServerAndDB2/storage"
	"github.com/sirupsen/logrus"
)

//Base API server description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	// Добавление поля для работы с хранилищем
	storage *storage.Storage
}

//API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start http server/configure loggers,router, db connection etc..
func (api *API) Start() error {
	// trying to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	// Подтверждение что логгер запустился
	api.logger.Info("starting api server at port:", api.config.BindAddr)
	//Конфигурируем  маршрутизатор
	api.configureRouterField()

	// Конфигурируем хранилище
	if err := api.configureStorageField(); err != nil {
		return err
	}
	//На этапе валидного завершения стартуем http-server

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
