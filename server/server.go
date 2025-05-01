// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/osyah/pletyvo-node/controller"
)

const shutdownTimeout = time.Second * 5

type Engine struct {
	notify chan error
	config Config
	server *fiber.App
}

func New(config Config, service *Service) *Engine {
	server := fiber.New(config.Fiber.Wrap())

	server.Use(cors.New(cors.Config{
		AllowOrigins:     config.CORS.AllowOrigins(),
		AllowMethods:     config.CORS.AllowMethods(),
		AllowHeaders:     config.CORS.AllowHeaders(),
		AllowCredentials: config.CORS.Credentials,
	}))

	server.Use(controller.NetworkMiddleware)

	api := server.Group("/api")
	{
		dapp := api.Group("/dapp")
		{
			controller.NewDAppV1(service.DApp).RegisterRoutes(dapp)
		}
		delivery := api.Group("/delivery")
		{
			controller.NewDeliveryV1(service.Delivery).RegisterRoutes(delivery)
		}
	}

	return &Engine{
		notify: make(chan error, 1),
		config: config,
		server: server,
	}
}

func (e Engine) Listen() {
	e.notify <- e.server.Listen(e.config.Address)
}

func (e Engine) Notify() <-chan error { return e.notify }

func (e Engine) Shutdown() error {
	return e.server.ShutdownWithTimeout(shutdownTimeout)
}
