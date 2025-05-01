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

type DeliveryChannel struct{ query delivery.ChannelQuery }

func NewDeliveryChannel(query delivery.ChannelQuery) *DeliveryChannel {
	return &DeliveryChannel{query: query}
}

func (dc DeliveryChannel) GetByID(ctx context.Context, id uuid.UUID) (*delivery.Channel, error) {
	return dc.query.GetByID(ctx, id)
}

func (DeliveryChannel) Create(context.Context, *delivery.ChannelCreateInput) (*dapp.EventResponse, error) {
	return nil, pletyvo.CodeNotImplemented
}

func (DeliveryChannel) Update(context.Context, *delivery.ChannelUpdateInput) (*dapp.EventResponse, error) {
	return nil, pletyvo.CodeNotImplemented
}
