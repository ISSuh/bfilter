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

package bfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		// given
		b := NewBitSet(8)

		// when
		b.set(0).set(2).set(4).set(6)

		// then
		target := uint8(0xAA)
		assert.Equal(t, target, b.bits[0])
	})

	t.Run("extendSet", func(t *testing.T) {
		// given
		b := NewBitSet(8)

		// when
		b.set(0).set(2).set(4).set(6).set(8)

		// then
		target := uint8(0xAA)
		assert.Equal(t, target, b.bits[0])
		assert.Equal(t, 2, len(b.bits))
	})
}

func TestGet(t *testing.T) {
	t.Run("getFromUnderOneByte", func(t *testing.T) {
		// given
		b := NewBitSet(4)

		// when
		b.set(0).set(2)

		// then
		assert.True(t, b.get(0))
		assert.False(t, b.get(1))
		assert.True(t, b.get(2))
		assert.False(t, b.get(3))
	})

	t.Run("getFromOverOneByte", func(t *testing.T) {
		// given
		b := NewBitSet(9)

		// when
		b.set(0).set(2).set(5).set(9)

		// then
		assert.True(t, b.get(0))
		assert.True(t, b.get(2))
		assert.True(t, b.get(5))
		assert.True(t, b.get(9))
		assert.False(t, b.get(10))
	})
}

func TestClear(t *testing.T) {
	t.Run("clear", func(t *testing.T) {
		// given
		b := NewBitSet(4)
		b.set(0).set(2)

		// when
		b.clear()

		// then
		assert.False(t, b.get(0))
		assert.False(t, b.get(2))
	})
}

func TestBytes(t *testing.T) {
	t.Run("bytes", func(t *testing.T) {
		// given
		b := NewBitSet(4)
		b.set(0).set(2)

		// when
		value := b.bytes()

		// then
		target := uint8(0xA0)
		assert.Equal(t, 1, len(value))
		assert.Equal(t, target, value[0])
	})
}

func BenchmarkBitset(b *testing.B) {
	size := 1000
	bitset := NewBitSet(size)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		location := uint(i % size)
		b.StartTimer()

		bitset.set(location)
	}
}

func BenchmarkBitsetSizePowerOfByte(b *testing.B) {
	size := 1024
	bitset := NewBitSet(size)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		location := uint(i % size)
		b.StartTimer()

		bitset.set(location)
	}
}

func BenchmarkGet(b *testing.B) {
	size := 1024
	bitset := NewBitSet(size)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bitset.set(uint(i))
		b.StartTimer()

		bitset.get(uint(i))
	}
}
