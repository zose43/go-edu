package intset

import (
	"bytes"
	"fmt"
)

type Intset struct {
	words []uint64
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

//func (s *Intset) IntersectWith(t *Intset) *Intset {
//	intrsc := new(Intset)
//	for i, word := range t.words {
//		// todo realize
//	}
//}

func (s *Intset) Elems() []int {
	return []int{}
	// todo realize
}

func (s *Intset) AddAll(items ...int) {
	for _, i := range items {
		s.Add(i)
	}
}

func (s *Intset) Remove(x int) {
	word, bit := x/64, uint64(x%64)
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
	buf.WriteString("{ ")
	for i, word := range s.words {
		if elems, ok := s.convert(i, word); ok {
			for _, elem := range elems {
				buf.WriteString(fmt.Sprintf("%d ", elem))
			}
		}
	}
	buf.WriteString(" }")
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
	word, bit := x/64, uint64(x%64)
	for word >= s.Len() {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *Intset) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < s.Len() && s.words[word]&(1<<bit) != 0
}
