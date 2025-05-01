// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"strings"

	"github.com/VictoriaMetrics/easyproto"
	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
	"github.com/osyah/go-pletyvo/delivery"
	"github.com/osyah/hryzun"
)

func (DeliveryPost) marshal(event *dapp.SystemEvent, input *delivery.PostInput) []byte {
	mr := mp.Get()

	mm := mr.MessageMarshaler()
	mm.AppendBytes(1, event.Hash[:])
	mm.AppendBytes(2, event.Author[:])
	mm.AppendString(3, input.Content)

	dst := mr.Marshal(nil)

	mp.Put(mr)

	return dst
}

func (DeliveryPost) unmarshal(src []byte, post *delivery.Post) (err error) {
	var fc easyproto.FieldContext

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return hryzun.NewStatus(
				pletyvo.CodeInternal,
				"pletyvo-node/store/pebble: cannot read next DeliveryPost field",
			)
		}

		switch fc.FieldNum {
		case 1:
			hash, ok := fc.Bytes()
			if !ok {
				return hryzun.NewStatus(
					pletyvo.CodeInternal,
					"pletyvo-node/store/pebble: cannot read DeliveryPost hash",
				)
			}

			post.Hash = dapp.Hash(hash)
		case 2:
			author, ok := fc.Bytes()
			if !ok {
				return hryzun.NewStatus(
					pletyvo.CodeInternal,
					"pletyvo-node/store/pebble: cannot read DeliveryPost author",
				)
			}

			post.Author = dapp.Hash(author)
		case 3:
			content, ok := fc.String()
			if !ok {
				return hryzun.NewStatus(
					pletyvo.CodeInternal,
					"pletyvo-node/store/pebble: cannot read DeliveryPost content",
				)
			}

			post.Content = strings.Clone(content)
		}
	}

	return
}
