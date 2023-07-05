package memo

type Memo struct {
	f     Fmemo
	cache map[string]result
}

func (m *Memo) Get(key string) (interface{}, error) {
	data, ok := m.cache[key]
	if !ok {
		data.value, data.err = m.f(key)
		m.cache[key] = data
	}
	return data.value, data.err
}

func New(f Fmemo) *Memo {
	return &Memo{f, make(map[string]result)}
}

type Fmemo func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}
