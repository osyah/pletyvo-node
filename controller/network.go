// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/osyah/go-pletyvo"
)

func NetworkMiddleware(ctx *fiber.Ctx) error {
	s := ctx.Get("Network")

	if len(s) != 0 {
		network, err := pletyvo.NetworkFromString(s)
		if err != nil {
			return ErrorHandler(ctx, err)
		}

		ctx.Context().SetUserValue(pletyvo.ContextKeyNetwork, network)
	}

	return ctx.Next()
}
