// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/hryzun"
)

func marshalDAppEvent(event *dapp.SystemEvent) []byte {
	m := mp.Get()

	mm := m.MessageMarshaler()
	mm.AppendBytes(1, event.Body)

	am := mm.AppendMessage(2)
	am.AppendUint32(1, uint32(event.Auth.Schema))
	am.AppendBytes(2, event.Auth.PublicKey)
	am.AppendBytes(3, event.Auth.Signature)

	dst := m.Marshal(nil)

	mp.Put(m)

	return dst
}

func (DAppEvent) unmarshal(src []byte, event *dapp.Event) (err error) {
	var fc easyproto.FieldContext

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return hryzun.NewStatus(
				pletyvo.CodeInternal,
				"pletyvo-node/store/pebble: cannot read next DAppEvent field",
			)
		}

		switch fc.FieldNum {
		case 1:
			body, ok := fc.Bytes()
			if !ok {
				return hryzun.NewStatus(
					pletyvo.CodeInternal,
					"pletyvo-node/store/pebble: cannot read DAppEvent body",
				)
			}

			event.Body = append(event.Body, body...)
		case 2:
			data, ok := fc.MessageData()
			if !ok {
				return hryzun.NewStatus(
					pletyvo.CodeInternal,
					"pletyvo-node/store/pebble: cannot read DAppEvent auth",
				)
			}

			if err = unmarshalDAppAuth(data, &event.Auth); err != nil {
				return err
			}
		}
	}

	return
}

func unmarshalDAppAuth(src []byte, auth *dapp.AuthHeader) (err error) {
	var fc easyproto.FieldContext

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return hryzun.NewStatus(
				pletyvo.CodeInternal,
				"pletyvo-node/store/pebble: cannot read next DApp auth field",
			)
		}

		switch fc.FieldNum {
		case 1:
			schema, ok := fc.Uint32()
			if !ok {
				return hryzun.NewStatus(
					pletyvo.CodeInternal,
					"pletyvo-node/store/pebble: cannot read DApp auth.schema",
				)
			}

			auth.Schema = byte(schema)
		case 2:
			publicKey, ok := fc.Bytes()
			if !ok {
				return hryzun.NewStatus(
					pletyvo.CodeInternal,
					"pletyvo-node/store/dapppebble: cannot read DApp auth.public_key",
				)
			}

			auth.PublicKey = append(auth.PublicKey, publicKey...)
		case 3:
			signature, ok := fc.Bytes()
			if !ok {
				return hryzun.NewStatus(
					pletyvo.CodeInternal,
					"pletyvo-node/store/pebble: cannot read DApp auth.signature",
				)
			}

			auth.Signature = append(auth.Signature, signature...)
		}
	}

	return
}
