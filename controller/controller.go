// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/go-pletyvo/delivery"
)

type DAppV1 struct{ service *dapp.Service }

func NewDAppV1(service *dapp.Service) *DAppV1 {
	return &DAppV1{service: service}
}

func (da DAppV1) RegisterRoutes(router fiber.Router) {
	v1 := router.Group("/v1")
	{
		NewDAppV1Event(da.service.Event).RegisterRoutes(v1)
		NewDAppV1Hash(da.service.Hash).RegisterRoutes(v1)
	}
}

type DeliveryV1 struct{ service *delivery.Service }

func NewDeliveryV1(service *delivery.Service) *DeliveryV1 {
	return &DeliveryV1{service: service}
}

func (d DeliveryV1) RegisterRoutes(router fiber.Router) {
	v1 := router.Group("/v1")
	{
		NewDeliveryV1Channel(d.service.Channel).RegisterRoutes(v1)
		NewDeliveryV1Message(d.service.Message).RegisterRoutes(v1)
		NewDeliveryV1Post(d.service.Post).RegisterRoutes(v1)
	}
}
