// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/go-pletyvo/delivery"
)

type DeliveryV1Message struct{ service delivery.MessageService }

func NewDeliveryV1Message(service delivery.MessageService) *DeliveryV1Message {
	return &DeliveryV1Message{service: service}
}

func (dm DeliveryV1Message) RegisterRoutes(router fiber.Router) {
	channel := router.Group("/channel")
	{
		channel.Post("/send", dm.sendHandler)
	}

	message := channel.Group("/:channel/messages")
	{
		message.Get("/", dm.getHandler)
		message.Get("/:message", dm.getByIDHandler)
	}
}

func (dm DeliveryV1Message) sendHandler(ctx *fiber.Ctx) error {
	var input dapp.EventInput

	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = dm.service.Send(ctx.Context(), &input)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (dm DeliveryV1Message) getHandler(ctx *fiber.Ctx) error {
	channel, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	option, err := QueryOption(ctx)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	posts, err := dm.service.Get(ctx.Context(), channel, option)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(posts)
}

func (dm DeliveryV1Message) getByIDHandler(ctx *fiber.Ctx) error {
	channel, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(ctx.Params("message"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	post, err := dm.service.GetByID(ctx.Context(), channel, id)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(post)
}
