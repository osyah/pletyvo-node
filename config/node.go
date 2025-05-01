// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package config

import (
	"github.com/osyah/pletyvo-node/server"
	"github.com/osyah/pletyvo-node/store/pebble"
)

type Node struct {
	Store  Store         `cfg:"store"`
	Server server.Config `cfg:"server"`
}

type Store struct {
	Pebble pebble.Config `cfg:"pebble"`
}
