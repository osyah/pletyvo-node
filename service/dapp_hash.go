// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package service

import (
	"context"

	"github.com/osyah/go-pletyvo/dapp"
)

type DAppHash struct{ query dapp.HashQuery }

func NewDAppHash(query dapp.HashQuery) *DAppHash {
	return &DAppHash{query: query}
}

func (dah DAppHash) GetByID(ctx context.Context, id dapp.Hash) (*dapp.EventResponse, error) {
	return dah.query.GetByID(ctx, id)
}
