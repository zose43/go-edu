package intset

import (
	"bytes"
	"fmt"
)

const bitCount = 32 << (^uint(0) >> 63)

type Intset struct {
	words []uint64
}

func (s *Intset) tokenize(x int) (int, uint64) {
	return x / bitCount, uint64(x % bitCount)
}

func (s *Intset) SymmetricDifference(t *Intset) *Intset {
	symmDiff := new(Intset)
	xdiff := s.DifferenceWith(t)
	ydiff := t.DifferenceWith(s)
	symmDiff.UnionWith(xdiff)
	symmDiff.UnionWith(ydiff)
	return symmDiff
}

func (s *Intset) convert(index int, word uint64) ([]int, bool) {
	var res []int
	if word == 0 {
		return res, false
	}
	for j := 0; j < 64; j++ {
		if word&(1<<uint(j)) != 0 {
			res = append(res, 64*index+j)
		}
	}
	return res, true
}

func (s *Intset) DifferenceWith(t *Intset) *Intset {
	diff := new(Intset)
	for i, word := range t.words {
		if elems, ok := t.convert(i, word); ok {
			for _, elem := range elems {
				if !s.Has(elem) {
					diff.Add(elem)
				}
			}
		}
	}
	return diff
}

func (s *Intset) IntersectWith(t *Intset) *Intset {
	intrsc := new(Intset)
	for i, word := range t.words {
		if elems, ok := t.convert(i, word); ok {
			for _, elem := range elems {
				if s.Has(elem) {
					intrsc.Add(elem)
				}
			}
		}
	}
	return intrsc
}

func (s *Intset) Elems() []int {
	var res []int
	for i, word := range s.words {
		if elems, ok := s.convert(i, word); ok {
			for _, elem := range elems {
				res = append(res, elem)
			}
		}
	}
	return res
}

func (s *Intset) AddAll(items ...int) {
	for _, i := range items {
		s.Add(i)
	}
}

func (s *Intset) Remove(x int) {
	word, bit := s.tokenize(x)
	if s.Has(x) {
		s.words[word] ^= 1 << uint(bit)
	}
}

func (s *Intset) Clear() {
	s.words = nil
}

func (s *Intset) Copy() *Intset {
	bitset := new(Intset)
	bitset.words = s.words
	return bitset
}

func (s *Intset) Len() int {
	return len(s.words)
}

func (s *Intset) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if elems, ok := s.convert(i, word); ok {
			for i, elem := range elems {
				if i == len(elems)-1 {
					buf.WriteString(fmt.Sprintf("%d", elem))
				} else {
					buf.WriteString(fmt.Sprintf("%d ", elem))
				}
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *Intset) UnionWith(t *Intset) {
	for i, tword := range t.words {
		if i < s.Len() {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *Intset) Add(x int) {
	word, bit := s.tokenize(x)
	for word >= s.Len() {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *Intset) Has(x int) bool {
	word, bit := s.tokenize(x)
	return word < s.Len() && s.words[word]&(1<<bit) != 0
}
