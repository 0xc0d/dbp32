# dbp32
Unsigned Integer 32 Byte Packing Compression. Inspired by [lemire/FastPFor](https://github.com/lemire/FastPFor).


Package bp32 is an implementation of the binary packing integer compression algorithm in in Go (also known as PackedBinary) using  unsigned 32 integer blocks.
It is mostly suitable for arrays containing small positive integers like IPv4 addresses or timestamp.
Given a list of sorted integers, it first compute the successive differences prior to compression.
For details, please see [Daniel Lemire and Leonid Boytsov, Decoding billions of integers per second](http://arxiv.org/abs/1209.2137).


# Usage
[This](https://play.golang.org/p/Fj6qslfA2sP) is the link to the go playground although `time.Since` does not work there.
```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/0xc0d/dbp32"
)

func main() {
	input := make([]uint32, 1e6)
	randomUint32(input, 1e3)
	output := make([]uint32, 1e6)

	s := time.Now()
	n, err := dbp32.Compress(input, output)
	if err != nil {
		panic(err)
	}
	e := time.Since(s)

	fmt.Println("Compressed in:", e.String(), "Ratio: ", float64(1e6)/float64(n))

	s = time.Now()
	n, err = dbp32.Decompress(output[:n], input)
	if err != nil {
		panic(err)
	}
	if n != 1e6 {
		panic("decompression failed")
	}
	e = time.Since(s)

	fmt.Println("Decompressed in:", e.String(), "ns")
}

func randomUint32(in []uint32, maxDist uint32) {
	last := rand.Uint32() % (maxDist)
	for i := range in {
		in[i] = last + rand.Uint32()%(maxDist)
		last = in[i]
	}
}
```

# Benchmark
```
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkCompress
BenchmarkCompress/standard_1000
BenchmarkCompress/standard_1000-16         	  488390	        2488 ns/op	     896 B/op	       1 allocs/op
BenchmarkCompress/standard_100000
BenchmarkCompress/standard_100000-16       	    7816	      143335 ns/op	     256 B/op	       1 allocs/op
BenchmarkCompress/standard_1000000
BenchmarkCompress/standard_1000000-16      	     880	     1733248 ns/op	     512 B/op	       1 allocs/op
BenchmarkCompress/standard_100000000
BenchmarkCompress/standard_100000000-16    	       7	   161584835 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompress/scored_1000
BenchmarkCompress/scored_1000-16           	  602750	        1730 ns/op	     896 B/op	       1 allocs/op
BenchmarkCompress/scored_100000
BenchmarkCompress/scored_100000-16         	    8806	      120738 ns/op	     256 B/op	       1 allocs/op
BenchmarkCompress/scored_1000000
BenchmarkCompress/scored_1000000-16        	     980	     1244335 ns/op	     512 B/op	       1 allocs/op
BenchmarkCompress/scored_100000000
BenchmarkCompress/scored_100000000-16      	       8	   150365733 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecompress
BenchmarkDecompress/standard_1000
BenchmarkDecompress/standard_1000-16         	 1000000	      1158 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecompress/standard_100000
BenchmarkDecompress/standard_100000-16       	   18090	     66249 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecompress/standard_1000000
BenchmarkDecompress/standard_1000000-16      	    1636	    659632 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecompress/standard_100000000
BenchmarkDecompress/standard_100000000-16    	      16	  71385376 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecompress/scored_1000
BenchmarkDecompress/scored_1000-16           	 1367968	     869.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecompress/scored_100000
BenchmarkDecompress/scored_100000-16         	   18804	     63270 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecompress/scored_1000000
BenchmarkDecompress/scored_1000000-16        	    1640	    637747 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecompress/scored_100000000
BenchmarkDecompress/scored_100000000-16      	      18	  70347136 ns/op	       0 B/op	       0 allocs/op

```
