// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/cockroachdb/pebble"
	"github.com/osyah/hryzun/module"
)

var mp easyproto.MarshalerPool

type Config struct {
	Dirname string `cfg:"dirname"`
}

func New(config Config) (*pebble.DB, error) {
	return pebble.Open(config.Dirname, &pebble.Options{})
}

func Register(base *module.Container, config Config) {
	base.RegisterHandler("pebble", func(base *module.Container) any {
		db, err := New(config)
		if err != nil {
			panic("pletyvo-node/store/pebble: " + err.Error())
		}

		base.RegisterCloser("pebble", db.Close)

		return db
	})

	base.RegisterHandler("dapp_pebble", func(base *module.Container) any {
		return NewDApp(module.Get[*pebble.DB](base, "pebble"))
	})
	base.RegisterHandler("delivery_pebble", func(base *module.Container) any {
		return NewDelivery(module.Get[*pebble.DB](base, "pebble"))
	})
}
