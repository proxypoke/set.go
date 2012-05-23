package set_test

type Set map[interface{}]bool

func (s Set) Add(x interface{}) (ret bool) {
	ret = s.Contains(x)
	s[x] = true
	return
}

func (s Set) Discard(x interface{}) (ret bool) {
	ret = s.Contains(x)
	s[x] = false
	return
}

func (s Set) Contains(x interface{}) (ret bool) {
	_, ret = s[x]
	return
}

func (s Set) Members() (mem []interface{}) {
	for m := range s {
		mem = append(mem, m)
	}
	return
}
