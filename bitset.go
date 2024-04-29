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

const (
	ByteSize        uint64 = 8
	PowerOfByteSize uint64 = 3
)

type bitset struct {
	bits []byte
}

func NewBitSet(size uint64) bitset {
	bitsetSize := (size >> PowerOfByteSize) + 1
	return bitset{
		bits: make([]byte, bitsetSize),
	}
}

func (b *bitset) set(location uint64) *bitset {
	index := location >> PowerOfByteSize
	offset := location % ByteSize

	bits := b.bits[index]

	// calculate bit and set
	target := byte((1 << (ByteSize - offset - 1)))
	bits |= target

	// set bits
	b.bits[index] = bits
	return b
}

func (b *bitset) get(location uint64) bool {
	index := location >> PowerOfByteSize
	offset := location % ByteSize

	if int(index) >= len(b.bits) {
		return false
	}

	bits := b.bits[index]
	target := byte((1 << (ByteSize - offset - 1)))
	return (bits & target) != 0
}

func (b *bitset) clear() {
	b.bits = b.bits[:0]
}

func (b *bitset) bytes() []byte {
	return b.bits
}
