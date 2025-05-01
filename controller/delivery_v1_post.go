// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/osyah/go-pletyvo/delivery"
)

type DeliveryV1Post struct{ service delivery.PostQuery }

func NewDeliveryV1Post(service delivery.PostQuery) *DeliveryV1Post {
	return &DeliveryV1Post{service: service}
}

func (dp DeliveryV1Post) RegisterRoutes(router fiber.Router) {
	post := router.Group("/channel/:channel/posts")
	{
		post.Get("/", dp.getHandler)
		post.Get("/:post", dp.getByIDHandler)
	}
}

func (dp DeliveryV1Post) getHandler(ctx *fiber.Ctx) error {
	channel, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	option, err := QueryOption(ctx)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	posts, err := dp.service.Get(ctx.Context(), channel, option)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(posts)
}

func (dp DeliveryV1Post) getByIDHandler(ctx *fiber.Ctx) error {
	channel, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(ctx.Params("post"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	post, err := dp.service.GetByID(ctx.Context(), channel, id)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(post)
}
