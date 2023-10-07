package utils

import "github.com/henrylee2cn/goutil/bitset"

const (
	default_size = 2 << 24
)

var seeds = [6]int{324, 235, 346, 563, 143, 654}

type BloomFilter struct {
	bitmap    *bitset.BitSet
	hashFuncs [6]SimpleHash
}

type SimpleHash struct {
	cap  int
	seed int
}

func NewBloomFilter() *BloomFilter {
	BF := new(BloomFilter)
	for i := 0; i < len(BF.hashFuncs); i++ {
		BF.hashFuncs[i] = SimpleHash{default_size, seeds[i]}
	}
	BF.bitmap = bitset.New(default_size)
	return BF
}

func (s *SimpleHash) Hash(val string) int {
	result := 0
	mask := s.cap - 1
	for _, v := range val {
		result += (result*s.seed)&mask + int(v)
	}
	return mask & result
}

func (b *BloomFilter) Insert(val string) {
	for _, curFunc := range b.hashFuncs {
		b.bitmap.Set(curFunc.Hash(val), true)
	}
}

func (b *BloomFilter) IsContains(val string) bool {
	if val == "" {
		return false
	}
	for _, curFunc := range b.hashFuncs {
		if !b.bitmap.Get(curFunc.Hash(val)) {
			return false
		}
	}
	return true
}
