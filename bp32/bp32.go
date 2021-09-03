package bp32

import (
	"math/bits"

	"github.com/0xc0d/dbp32/bitpack"
)

const (
	defaultBlockSize = 128
	steps            = defaultBlockSize / 4
)

func Compress(in, out []uint32) (m, n int, err error) {
	m = len(in) - len(in)%defaultBlockSize
	if m == 0 {
		return
	}

	out[n] = uint32(m)
	n++

	var offset uint32
	for i := 0; i < m; i += defaultBlockSize {
		maxBits1 := maxBits(offset, in[i:i+steps])
		offset2 := in[i+31]
		maxBits2 := maxBits(offset2, in[i+steps:i+2*steps])
		offset3 := in[i+32+31]
		maxBits3 := maxBits(offset3, in[i+2*steps:i+3*steps])
		offset4 := in[i+2*32+31]
		maxBits4 := maxBits(offset4, in[i+3*steps:i+4*steps])

		out[n] = (maxBits1 << 24) | (maxBits2 << 16) | (maxBits3 << 8) | maxBits4
		n++

		if err = bitpack.DeltaPack(offset, in[i:], out[n:], int(maxBits1)); err != nil {
			return
		}
		n += int(maxBits1)

		if err = bitpack.DeltaPack(offset2, in[i+steps:], out[n:], int(maxBits2)); err != nil {
			return
		}
		n += int(maxBits2)

		if err = bitpack.DeltaPack(offset3, in[i+2*steps:], out[n:], int(maxBits3)); err != nil {
			return
		}
		n += int(maxBits3)

		if err = bitpack.DeltaPack(offset4, in[i+3*steps:], out[n:], int(maxBits4)); err != nil {
			return
		}
		n += int(maxBits4)

		offset = in[i+4*steps-1]
	}

	return
}

func Decompress(in, out []uint32) (m, n int, err error) {
	length := len(in)
	if length == 0 {
		return
	}

	n = int(in[m])
	m++

	var offset uint32
	for s := 0; s < n; s += defaultBlockSize {
		tmp := in[m]
		maxBits1 := tmp >> 24
		maxBits2 := (tmp >> 16) & 0xFF
		maxBits3 := (tmp >> 8) & 0xFF
		maxBits4 := (tmp) & 0xFF

		m++

		if err = bitpack.DeltaUnpack(offset, in[m:], out[s:], int(maxBits1)); err != nil {
			return
		}
		m += int(maxBits1)
		offset = out[s+steps-1]

		if err = bitpack.DeltaUnpack(offset, in[m:], out[s+steps:], int(maxBits2)); err != nil {
			return
		}
		m += int(maxBits2)
		offset = out[s+2*steps-1]

		if err = bitpack.DeltaUnpack(offset, in[m:], out[s+2*steps:], int(maxBits3)); err != nil {
			return
		}
		m += int(maxBits3)
		offset = out[s+3*steps-1]

		if err = bitpack.DeltaUnpack(offset, in[m:], out[s+3*steps:], int(maxBits4)); err != nil {
			return
		}
		m += int(maxBits4)
		offset = out[s+4*steps-1]
	}

	return
}

func maxBits(offset uint32, buf []uint32) uint32 {
	var mask uint32
	for _, v := range buf {
		mask |= v - offset
		offset = v
	}

	return uint32(bits.Len32(mask))
}
