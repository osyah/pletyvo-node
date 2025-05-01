// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/go-pletyvo/delivery"
)

type DeliveryPost struct{ query delivery.PostQuery }

func NewDeliveryPost(query delivery.PostQuery) *DeliveryPost {
	return &DeliveryPost{query: query}
}

func (dp DeliveryPost) Get(ctx context.Context, id uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Post, error) {
	option.Prepare()

	return dp.query.Get(ctx, id, option)
}

func (dp DeliveryPost) GetByID(ctx context.Context, channel, id uuid.UUID) (*delivery.Post, error) {
	return dp.query.GetByID(ctx, channel, id)
}

func (DeliveryPost) Create(context.Context, *delivery.PostCreateInput) (*dapp.EventResponse, error) {
	return nil, pletyvo.CodeNotImplemented
}

func (DeliveryPost) Update(context.Context, *delivery.PostUpdateInput) (*dapp.EventResponse, error) {
	return nil, pletyvo.CodeNotImplemented
}
