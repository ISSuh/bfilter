# bloom

implement bloom filter on Golang

### useage

```bash
go get github.com/ISSuh/bloom
```

```go

import (
  "hash/fnv"

  "github.com/ISSuh/bloom"
)

func ExampleDefaultHash() {
  numberOfSize := uint64(1024)
  numberOfHash := 3

  // use default hash function(murmur3)
  fiter := bloom.NewFilter(numberOfSize, numberOfHash)

  testKey := []byte("test")
  if err := filter.Add(testKey); err != nil {
    return
  }

  if has, err := filter.Has(testKey); !has || err != nil {
    return
  }
}

func ExampleUserHash() {
  numberOfSize := uint64(1024)
  numberOfHash := 3
  userHash := fnv.New64()

  // use user hash function
  fiter := bloom.NewFilterWithHash(numberOfSize, numberOfHash, userHash)

  testKey := []byte("test")
  if err := filter.Add(testKey); err != nil {
    return
  }

  if has, err := filter.Has(testKey); !has || err != nil {
    return
  }
}

```

### benchmark

```bash
> go test -bench . -benchmem -benchtime 10000x
goos: darwin
goarch: arm64
pkg: github.com/ISSuh/bloom
BenchmarkBitsetSet-11                              10000                30.33 ns/op            0 B/op          0 allocs/op
BenchmarkBitsetSetSizePowerOfByte-11               10000                26.44 ns/op            0 B/op          0 allocs/op
BenchmarkBitsetGet-11                              10000                21.33 ns/op            0 B/op          0 allocs/op
BenchmarkFilterAdd-11                              10000                98.44 ns/op            0 B/op          0 allocs/op
BenchmarkFilterHas-11                              10000                87.97 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/ISSuh/bloom  2.238s
```
