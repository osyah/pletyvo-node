// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Address string      `cfg:"address"`
	Fiber   FiberConfig `cfg:"fiber"`
	CORS    CORS        `cfg:"cors"`
}

type FiberConfig struct {
	Prefork     bool   `cfg:"prefork"`
	BodyLimit   int    `cfg:"body_limit"`
	Concurrency int    `cfg:"concurrency"`
	AppName     string `cfg:"app_name"`
}

func (fc FiberConfig) Wrap() fiber.Config {
	return fiber.Config{
		Prefork:     fc.Prefork,
		BodyLimit:   fc.BodyLimit,
		Concurrency: fc.Concurrency,
		AppName:     fc.AppName,
	}
}

type CORS struct {
	Origins     []string `cfg:"origins"`
	Methods     []string `cfg:"methods"`
	Headers     []string `cfg:"headers"`
	Credentials bool     `cfg:"credentials"`
}

func (c CORS) AllowOrigins() string {
	return strings.Join(c.Origins, ",")
}

func (c CORS) AllowMethods() string {
	return strings.Join(c.Methods, ",")
}

func (c CORS) AllowHeaders() string {
	return strings.Join(c.Headers, ",")
}
