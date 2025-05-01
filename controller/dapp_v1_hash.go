// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/osyah/go-pletyvo/dapp"
)

type DAppV1Hash struct{ service dapp.HashQuery }

func NewDAppV1Hash(service dapp.HashQuery) *DAppV1Hash {
	return &DAppV1Hash{service: service}
}

func (dah DAppV1Hash) RegisterRoutes(router fiber.Router) {
	router.Get("/hash/:id", dah.getByIDHandler)
}

func (dah DAppV1Hash) getByIDHandler(ctx *fiber.Ctx) error {
	id, err := dapp.HashFromString(ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	header, err := dah.service.GetByID(ctx.Context(), id)
	if err != nil {
		return ErrorHandler(ctx, err)
	}

	return ctx.JSON(header)
}
