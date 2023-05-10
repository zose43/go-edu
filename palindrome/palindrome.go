package palindrome

import "sort"

type Palindrome []rune

func (p *Palindrome) Len() int {
	return len(*p)
}

func (p *Palindrome) Less(i, j int) bool {
	return (*p)[i] < (*p)[j]
}

func (p *Palindrome) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *Palindrome) IsPalindrome(s sort.Interface) bool {
	sort.Sort(s)
	for i := 0; i < len(*p)-1; i += 2 {
		if res := !s.Less(i+1, i) && !s.Less(i, i+1); !res {
			return false
		}
	}
	return true
}
