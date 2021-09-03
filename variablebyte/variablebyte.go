package variablebyte

import (
	"bytes"
	"encoding/binary"
)

func Compress(in []uint32, out []uint32) (n int, err error) {
	length := len(in)
	if length == 0 {
		return
	}

	var buf bytes.Buffer
	buf.Grow(length * 8)
	var offset uint32

	for _, v := range in {
		val := v - offset
		offset = v

		for val >= 0x80 {
			buf.WriteByte(byte(val) | 0x80)
			val >>= 7
		}
		buf.WriteByte(byte(val))
	}

	for buf.Len()%4 != 0 {
		buf.WriteByte(0x80)
	}

	n = buf.Len() / 4
	for i := 0; i < n; i++ {
		out[i] = binary.BigEndian.Uint32(buf.Next(4))
	}

	return
}

func Decompress(in []uint32, out []uint32) (n int, err error) {
	length := len(in)
	if length == 0 {
		return
	}

	var s, p, shift int
	var v, offset uint32

	for p < length {
		c := in[p] >> (24 - s)
		s += 8

		if s == 32 {
			s = 0
			p++
		}

		v += (c & 127) << shift
		if c&128 == 0 {
			out[n] = v + offset
			offset += v
			n++
			v = 0
			shift = 0
		} else {
			shift += 7
		}

	}

	return
}
