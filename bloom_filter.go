// MIT License

// Copyright (c) 2024 ISSuh

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package bloom

import (
	"encoding/binary"
	"hash"

	"github.com/spaolacci/murmur3"
)

type Filter struct {
	// number of bitset size
	m uint64

	// number of hash
	k int

	// bitset
	v bitset

	// hash function
	// default hash function is murmur3
	h hash.Hash64
}

func NewFilter(size uint64, numberOfHash int) *Filter {
	return &Filter{
		m: size,
		k: numberOfHash,
		v: NewBitSet(size),
		h: murmur3.New64(),
	}
}

func NewFilterWithHash(size uint64, numberOfHash int, userHash hash.Hash64) *Filter {
	return &Filter{
		m: size,
		k: numberOfHash,
		v: NewBitSet(size),
		h: userHash,
	}
}

func (f *Filter) Add(key []byte) error {
	for i := uint32(0); i < uint32(f.k); i++ {
		index, err := f.location(key, i)
		if err != nil {
			return err
		}

		f.v.set(index)
	}
	return nil
}

func (f *Filter) Has(key []byte) (bool, error) {
	for i := uint32(0); i < uint32(f.k); i++ {
		if index, err := f.location(key, i); err != nil || !f.v.get(index) {
			return false, nil
		}
	}
	return true, nil
}

func (f *Filter) location(key []byte, seed uint32) (uint64, error) {
	f.h.Reset()

	if _, err := f.h.Write(key); err != nil {
		return 0, nil
	}

	s := f.h.Sum64()
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, s)

	lower := binary.BigEndian.Uint32(b[4:])
	higher := binary.BigEndian.Uint32(b[:4])

	v := (uint64(lower) + (uint64(higher) * uint64(seed))) % f.m
	return v, nil
}
