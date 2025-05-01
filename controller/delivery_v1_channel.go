// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/osyah/go-pletyvo/delivery"
)

type DeliveryV1Channel struct{ service delivery.ChannelQuery }

func NewDeliveryV1Channel(service delivery.ChannelQuery) *DeliveryV1Channel {
	return &DeliveryV1Channel{service: service}
}

func (dc DeliveryV1Channel) RegisterRoutes(router fiber.Router) {
	channel := router.Group("/channel/:channel")
	{
		channel.Get("/", dc.getByIDHandler)
	}
}

func (dc DeliveryV1Channel) getByIDHandler(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	channel, err := dc.service.GetByID(ctx.Context(), id)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(channel)
}
