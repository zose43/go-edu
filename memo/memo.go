package memo

import "os"

type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

type Fmemo func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

func (e *entry) deliver(resp chan<- result) {
	<-e.ready
	resp <- e.res
}

func (e *entry) call(key string, f Fmemo) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (m *Memo) Close() {
	close(m.requests)
}

func (m *Memo) Get(key string) (interface{}, error) {
	resp := make(chan result)
	m.requests <- request{key: key, response: resp}
	data := <-resp
	return data.value, data.err
}

func (m *Memo) server(f Fmemo, cancel <-chan struct{}) {
	cache := make(map[string]*entry)

	for {
		select {
		case <-cancel:
			for r := range m.requests {
				go func(r request) {
					val, err := f(r.key)
					r.response <- result{val, err}
				}(r)
			}
			return
		default:
			for r := range m.requests {
				e := cache[r.key]
				if e == nil {
					e = &entry{ready: make(chan struct{})}
					cache[r.key] = e
					go e.call(r.key, f)
				}
				go e.deliver(r.response)
			}
			return
		}
	}
}

func New(f Fmemo) *Memo {
	cancel := make(chan struct{})
	m := &Memo{make(chan request)}
	go m.server(f, cancel)

	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		cancel <- struct{}{}
	}()

	return m
}
