package main

import (
	"github.com/pkg/errors"
	"github.com/rtemb/srv-users/internal/config"
	"github.com/rtemb/srv-users/internal/handler"
	"github.com/rtemb/srv-users/internal/service"
	"github.com/rtemb/srv-users/internal/storage"
	"github.com/rtemb/srv-users/internal/token_encoder"
	"github.com/rtemb/srv-users/pkg/version"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New().WithFields(logrus.Fields{
		"gitSha":  version.GitSha,
		"version": version.ServiceVersion,
		"logger":  "cmd/srv-users",
	})
	logger.Println("loading service configurations")
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(errors.Wrap(err, "could not load service config"))
	}

	lvl, err := logrus.ParseLevel(cfg.AppConfig.LogLevel)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "could parse log level"))
	}
	logger.Logger.SetLevel(lvl)

	store := storage.NewStorage(cfg.Redis)
	tokenEncoder := token_encoder.NewEncoder(cfg.AppConfig.TokenKey)
	s := service.NewService(logger, store, tokenEncoder)
	grpcHandler := handler.NewHandler(s, logger)

	handler.StartService(cfg.Server, grpcHandler, logger)
}
