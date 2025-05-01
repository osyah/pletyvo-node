// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/osyah/go-pletyvo/dapp"
)

type DAppV1Event struct{ service dapp.EventService }

func NewDAppV1Event(service dapp.EventService) *DAppV1Event {
	return &DAppV1Event{service: service}
}

func (dae DAppV1Event) RegisterRoutes(router fiber.Router) {
	event := router.Group("/events")
	{
		event.Get("/", dae.getHandler)
		event.Post("/", dae.createHandler)
		event.Get("/:id", dae.getByIDHandler)
	}
}

func (dae DAppV1Event) getHandler(ctx *fiber.Ctx) error {
	option, err := QueryOption(ctx)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	events, err := dae.service.Get(ctx.Context(), option)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(events)
}

func (dae DAppV1Event) createHandler(ctx *fiber.Ctx) error {
	var input dapp.EventInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	response, err := dae.service.Create(ctx.Context(), &input)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(response)
}

func (dae DAppV1Event) getByIDHandler(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	event, err := dae.service.GetByID(ctx.Context(), id)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(event)
}
