package app

import (
	"context"
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/config"
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/http"
	"github.com/sirupsen/logrus"
)

// Application представляет основное приложение.
type Application struct {
	Config  *config.Config
	Handler *http.Handler
	log     *logrus.Logger
}

// New создает новый объект Application и настраивает его.
func New(ctx context.Context, log *logrus.Logger) (*Application, error) {
	// инициализирует конфигурацию
	cfg := config.MustLoad()
	log.WithField("config", cfg).Info("config parsed")

	h := http.NewHandler(cfg.BaseURL, log)
	// инициализирует объект Application
	app := &Application{
		Config:  cfg,
		Handler: h,
		log:     log,
	}

	return app, nil
}
