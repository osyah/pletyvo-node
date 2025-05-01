// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package service

import (
	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/go-pletyvo/delivery"
)

func NewDapp(repos *dapp.Repository, handler *dapp.Handler) *dapp.Service {
	return &dapp.Service{
		Event: NewDAppEvent(repos, handler), Hash: NewDAppHash(repos.Hash),
	}
}

func NewDelivery(repos *delivery.Repository) *delivery.Service {
	return &delivery.Service{
		Channel: NewDeliveryChannel(repos.Channel),
		Message: NewDeliveryMessage(repos.Message),
		Post:    NewDeliveryPost(repos.Post),
	}
}
