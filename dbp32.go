package dbp32

import (
	"github.com/0xc0d/dbp32/bp32"
	"github.com/0xc0d/dbp32/variablebyte"
)

func Compress(in []uint32, out []uint32) (n int, err error) {
	m, n, err := bp32.Compress(in, out)
	if err != nil {
		return
	}

	nn, err := variablebyte.Compress(in[m:], out[n:])
	n += nn

	return
}

func Decompress(in []uint32, out []uint32) (n int, err error) {
	m, n, err := bp32.Decompress(in, out)
	if err != nil {
		return
	}

	nn, err := variablebyte.Decompress(in[m:], out[n:])
	n += nn

	return
}
