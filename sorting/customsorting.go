package sorting

type StringSlice []string

func (sl *StringSlice) Len() int {
	return len(*sl)
}

func (sl *StringSlice) Less(i, j int) bool {
	return (*sl)[i] < (*sl)[j]
}

func (sl *StringSlice) Swap(i, j int) {
	(*sl)[i], (*sl)[j] = (*sl)[j], (*sl)[i]
}
