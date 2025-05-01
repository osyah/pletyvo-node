// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/go-pletyvo/delivery"
)

type DeliveryMessage struct{ repos delivery.MessageRepository }

func NewDeliveryMessage(repos delivery.MessageRepository) *DeliveryMessage {
	return &DeliveryMessage{repos: repos}
}

func (dm DeliveryMessage) Get(ctx context.Context, id uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Message, error) {
	option.Prepare()

	return dm.repos.Get(ctx, id, option)
}

func (dm DeliveryMessage) GetByID(ctx context.Context, channel uuid.UUID, id uuid.UUID) (*delivery.Message, error) {
	return dm.repos.GetByID(ctx, channel, id)
}

func (dm DeliveryMessage) Send(ctx context.Context, message *dapp.EventInput) error {
	if message.Body.Version() != dapp.EventBodyBasic {
		return dapp.ErrInvalidEventBodyVersion
	}

	if message.Body.Type() != delivery.MessageCreate {
		return dapp.ErrInvalidEventType
	}

	if !message.Verify(dapp.EventInputVerifier) {
		return pletyvo.CodeUnauthorized
	}

	var input delivery.MessageInput

	err := message.Body.Unmarshal(&input)
	if err != nil {
		return pletyvo.CodeInvalidArgument
	}

	if err = input.Validate(); err != nil {
		return err
	}

	if input.ID.Version() != 7 {
		return pletyvo.CodeInvalidArgument
	}

	sec, _ := input.ID.Time().UnixTime()
	interval := time.Now().Unix() - sec
	if interval > 5 || interval < -5 {
		return delivery.ErrInvalidMessageTime
	}

	return dm.repos.Create(ctx, message, &input)
}
