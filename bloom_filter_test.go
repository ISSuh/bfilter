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
	"hash/fnv"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestFilterSize uint64 = 1024
)

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		// given
		f := NewFilter(TestFilterSize, 3)

		// when
		testKey := []byte("test")
		err := f.Add(testKey)
		assert.Nil(t, err)

		// then
		has, err := f.Has(testKey)
		assert.Nil(t, err)
		assert.True(t, has)
	})

	fnv.New64()
}

func BenchmarkFilterAdd(b *testing.B) {
	f := NewFilter(TestFilterSize, 3)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		key := []byte(strconv.Itoa(i))
		b.StartTimer()

		f.Add(key)
	}
}

func BenchmarkFilterHas(b *testing.B) {
	f := NewFilter(TestFilterSize, 3)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		key := []byte(strconv.Itoa(i))
		f.Add(key)
		b.StartTimer()

		if has, err := f.Has(key); !has || err != nil {
			b.Fatal()
		}
	}
}
