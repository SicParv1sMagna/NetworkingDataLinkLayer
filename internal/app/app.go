package app

import (
	"context"
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/config"
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/http"
)

// Application представляет основное приложение.
type Application struct {
	Config  *config.Config
	Handler *http.Handler
}

// New создает новый объект Application и настраивает его.
func New(ctx context.Context) (*Application, error) {
	// инициализирует конфигурацию
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}
	h := http.NewHandler(cfg.BaseURL)
	// инициализирует объект Application
	app := &Application{
		Config:  cfg,
		Handler: h,
	}

	return app, nil
}
