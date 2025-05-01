// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/go-pletyvo/delivery"
	"github.com/rs/zerolog/log"

	"github.com/osyah/pletyvo-node/config"
	"github.com/osyah/pletyvo-node/server"
	"github.com/osyah/pletyvo-node/service"
	"github.com/osyah/pletyvo-node/store/pebble"
)

func Run(configPath string) {
	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	db, err := pebble.New(cfg.Node.Store.Pebble)
	if err != nil {
		log.Error().Err(err).Send()
	}

	handler := dapp.NewHandler()

	deliveryRepos := pebble.NewDelivery(db)
	delivery.NewExecutor(deliveryRepos).Register(handler)

	server := server.New(cfg.Node.Server, &server.Service{
		DApp:     service.NewDapp(pebble.NewDApp(db), handler),
		Delivery: service.NewDelivery(deliveryRepos),
	})

	go server.Listen()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Str("signal", s.String()).Send()
	case err := <-server.Notify():
		log.Error().Err(err).Send()
	}

	if err := server.Shutdown(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
