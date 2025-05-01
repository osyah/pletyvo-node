// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/cockroachdb/pebble"
	"github.com/osyah/go-pletyvo/dapp"
)

const (
	DAppEventPrefix = 1
	DAppHashPrefix  = 2
)

func NewDApp(db *pebble.DB) *dapp.Repository {
	return &dapp.Repository{
		Event: NewDAppEvent(db),
		Hash:  NewDAppHash(db),
	}
}
