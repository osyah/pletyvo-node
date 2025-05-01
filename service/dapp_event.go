// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
)

type DAppEvent struct {
	query   dapp.EventQuery
	hash    dapp.HashQuery
	handler *dapp.Handler
}

func NewDAppEvent(repos *dapp.Repository, handler *dapp.Handler) *DAppEvent {
	return &DAppEvent{query: repos.Event, hash: repos.Hash, handler: handler}
}

func (dae DAppEvent) Get(ctx context.Context, option *pletyvo.QueryOption) ([]*dapp.Event, error) {
	option.Prepare()

	return dae.query.Get(ctx, option)
}

func (dae DAppEvent) GetByID(ctx context.Context, id uuid.UUID) (*dapp.Event, error) {
	return dae.query.GetByID(ctx, id)
}

func (dae DAppEvent) Create(ctx context.Context, input *dapp.EventInput) (*dapp.EventResponse, error) {
	header := &dapp.EventHeader{
		Hash: dapp.NewHash(input.Auth.Schema, input.Auth.Signature),
	}

	response, err := dae.hash.GetByID(ctx, header.Hash)
	if err == nil {
		return response, nil
	}

	if !input.Verify(dapp.EventInputVerifier) {
		return nil, pletyvo.CodeUnauthorized
	}

	if input.Body.Version() == dapp.EventBodyLinked {
		_, err = dae.hash.GetByID(ctx, input.Body.Parent())
		if err != nil {
			return nil, pletyvo.CodeConflict
		}
	}

	header.ID, err = uuid.NewV7()
	if err != nil {
		return nil, pletyvo.CodeInternal
	}

	if err = dae.handler.Handle(ctx, &dapp.SystemEvent{
		EventHeader: header,
		EventInput:  input,
		Author:      dapp.NewHash(input.Auth.Schema, input.Auth.PublicKey),
	}); err != nil {
		return nil, pletyvo.CodeInternal
	}

	return &dapp.EventResponse{ID: header.ID}, nil
}
