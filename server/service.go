// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package server

import (
	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/go-pletyvo/delivery"
)

type Service struct {
	DApp     *dapp.Service
	Delivery *delivery.Service
}
