package dbp32

import (
	"math/rand"
	"testing"
)

func TestCompressAndDecompress(t *testing.T) {
	tests := []struct {
		name    string
		in      []uint32
		wantErr bool
	}{
		{
			name: "standard 1000",
			in: func() []uint32 {
				in := make([]uint32, 1e3)
				randomUint32(in, 1e3)
				return in
			}(),
			wantErr: false,
		},
		{
			name: "standard 100000000",
			in: func() []uint32 {
				in := make([]uint32, 1e8)
				randomUint32(in, 1e3)
				return in
			}(),
			wantErr: false,
		},
		{
			name: "scored 1000",
			in: func() []uint32 {
				in := make([]uint32, 1e3)
				randomUint32(in, 7)
				return in
			}(),
			wantErr: false,
		},
		{
			name: "scored 100000000",
			in: func() []uint32 {
				in := make([]uint32, 1e8)
				randomUint32(in, 7)
				return in
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make([]uint32, len(tt.in)*2)

			n, err := Compress(tt.in, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("%d -> %d \t ratio: %.2f", len(tt.in), n, float64(len(tt.in))/float64(n))

			final := make([]uint32, len(tt.in))
			n, err = Decompress(out[:n], final)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for i := range tt.in {
				if tt.in[i] != final[i] {
					t.Errorf("malformed at position %d input: %d, decompressed: %d", i, tt.in[i], final[i])
				}
			}

		})
	}
}

func BenchmarkCompress(b *testing.B) {
	tests := []struct {
		name string
		in   []uint32
	}{
		{
			name: "standard 1000",
			in: func() []uint32 {
				in := make([]uint32, 1e3)
				randomUint32(in, 1e3)
				return in
			}(),
		},
		{
			name: "standard 100000",
			in: func() []uint32 {
				in := make([]uint32, 1e5)
				randomUint32(in, 1e3)
				return in
			}(),
		},
		{
			name: "standard 1000000",
			in: func() []uint32 {
				in := make([]uint32, 1e6)
				randomUint32(in, 1e3)
				return in
			}(),
		},
		{
			name: "standard 100000000",
			in: func() []uint32 {
				in := make([]uint32, 1e8)
				randomUint32(in, 1e3)
				return in
			}(),
		},
		{
			name: "scored 1000",
			in: func() []uint32 {
				in := make([]uint32, 1e3)
				randomUint32(in, 7)
				return in
			}(),
		},
		{
			name: "scored 100000",
			in: func() []uint32 {
				in := make([]uint32, 1e5)
				randomUint32(in, 7)
				return in
			}(),
		},
		{
			name: "scored 1000000",
			in: func() []uint32 {
				in := make([]uint32, 1e6)
				randomUint32(in, 7)
				return in
			}(),
		},
		{
			name: "scored 100000000",
			in: func() []uint32 {
				in := make([]uint32, 1e8)
				randomUint32(in, 7)
				return in
			}(),
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			out := make([]uint32, len(tt.in)*2)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Compress(tt.in, out)
			}
		})
	}
}

func BenchmarkDecompress(b *testing.B) {
	tests := []struct {
		name string
		in   []uint32
	}{
		{
			name: "standard 1000",
			in: func() []uint32 {
				in := make([]uint32, 1e3)
				randomUint32(in, 1e3)
				return in
			}(),
		},
		{
			name: "standard 100000",
			in: func() []uint32 {
				in := make([]uint32, 1e5)
				randomUint32(in, 1e3)
				return in
			}(),
		},
		{
			name: "standard 1000000",
			in: func() []uint32 {
				in := make([]uint32, 1e6)
				randomUint32(in, 1e3)
				return in
			}(),
		},
		{
			name: "standard 100000000",
			in: func() []uint32 {
				in := make([]uint32, 1e8)
				randomUint32(in, 1e3)
				return in
			}(),
		},
		{
			name: "scored 1000",
			in: func() []uint32 {
				in := make([]uint32, 1e3)
				randomUint32(in, 7)
				return in
			}(),
		},
		{
			name: "scored 100000",
			in: func() []uint32 {
				in := make([]uint32, 1e5)
				randomUint32(in, 7)
				return in
			}(),
		},
		{
			name: "scored 1000000",
			in: func() []uint32 {
				in := make([]uint32, 1e6)
				randomUint32(in, 7)
				return in
			}(),
		},
		{
			name: "scored 100000000",
			in: func() []uint32 {
				in := make([]uint32, 1e8)
				randomUint32(in, 7)
				return in
			}(),
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			out := make([]uint32, len(tt.in)*2)
			n, _ := Compress(tt.in, out)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Decompress(out[:n], tt.in)
			}
		})
	}
}

func randomUint32(in []uint32, maxDist uint32) {
	last := rand.Uint32() % (maxDist)
	for i := range in {
		in[i] = last + rand.Uint32()%(maxDist)
		last = in[i]
	}
}
