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