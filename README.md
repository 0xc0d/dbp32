# dbp32
Unsigned Integer 32 Byte Packing Compression. Heavily forked from [zentures/encoding](https://github.com/zentures/encoding)


Package bp32 is an implementation of the binary packing integer compression algorithm in in Go (also known as PackedBinary) using  unsigned 32 integer blocks.
It is mostly suitable for arrays containing small positive integers like IPv4 addresses or timestamp.
Given a list of sorted integers, it first compute the successive differences prior to compression.
For details, please see [Daniel Lemire and Leonid Boytsov, Decoding billions of integers per second](http://arxiv.org/abs/1209.2137)


# Benchmark
```

```
