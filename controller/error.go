// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/osyah/go-pletyvo"
	"github.com/osyah/hryzun"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	switch t := err.(type) {
	case hryzun.Status:
		return ctx.Status(WrapStatus(t.Code)).
			JSON(&ErrorResponse{Message: t.Message})
	case hryzun.Code:
		return ctx.SendStatus(WrapStatus(t))
	default:
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
}

func WrapStatus(code hryzun.Code) int {
	switch code {
	case pletyvo.CodeInternal:
		return fiber.StatusInternalServerError
	case pletyvo.CodeNotFound:
		return fiber.StatusNotFound
	case pletyvo.CodePermissionDenied:
		return fiber.StatusForbidden
	case pletyvo.CodeInvalidArgument:
		return fiber.StatusBadRequest
	case pletyvo.CodeUnauthorized:
		return fiber.StatusUnauthorized
	case pletyvo.CodeConflict:
		return fiber.StatusConflict
	case pletyvo.CodeNotImplemented:
		return fiber.StatusNotImplemented
	default:
		return fiber.StatusBadRequest
	}
}
