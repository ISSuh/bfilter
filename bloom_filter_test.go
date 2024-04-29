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
	TestFilterHash int    = 3
)

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		// given
		f := NewFilter(TestFilterSize, TestFilterHash)

		// when
		testKey := []byte("test")
		err := f.Add(testKey)

		// then
		assert.Nil(t, err)
	})

	t.Run("userHash", func(t *testing.T) {
		// given
		userHash := fnv.New64()
		f, err := NewFilterWithHash(TestFilterSize, TestFilterHash, userHash)
		assert.Nil(t, err)

		// when
		testKey := []byte("test")
		err = f.Add(testKey)

		// then
		assert.Nil(t, err)
	})

	t.Run("emptyUserHash", func(t *testing.T) {
		// given
		// when
		f, err := NewFilterWithHash(TestFilterSize, TestFilterHash, nil)

		// then
		assert.NotNil(t, err)
		assert.Nil(t, f)
	})
}

func TestHas(t *testing.T) {
	t.Run("hasTrue", func(t *testing.T) {
		// given
		f := NewFilter(TestFilterSize, TestFilterHash)

		// when
		testKey := []byte("test")
		err := f.Add(testKey)
		assert.Nil(t, err)

		// then
		has, err := f.Has(testKey)
		assert.Nil(t, err)
		assert.True(t, has)
	})

	t.Run("hasFalse", func(t *testing.T) {
		// given
		f := NewFilter(TestFilterSize, TestFilterHash)

		// when
		testKey := []byte("test")
		has, err := f.Has(testKey)

		// then
		assert.Nil(t, err)
		assert.False(t, has)
	})

	t.Run("userHash", func(t *testing.T) {
		// given
		userHash := fnv.New64()
		f, err := NewFilterWithHash(TestFilterSize, TestFilterHash, userHash)
		assert.Nil(t, err)

		// when
		testKey := []byte("test")
		err = f.Add(testKey)
		assert.Nil(t, err)

		// then
		has, err := f.Has(testKey)
		assert.Nil(t, err)
		assert.True(t, has)
	})
}

func BenchmarkFilterAdd(b *testing.B) {
	f := NewFilter(TestFilterSize, TestFilterHash)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		key := []byte(strconv.Itoa(i))
		b.StartTimer()

		if err := f.Add(key); err != nil {
			b.Fatal()
		}
	}
}

func BenchmarkFilterHas(b *testing.B) {
	f := NewFilter(TestFilterSize, TestFilterHash)

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
