package set

type Set map[interface{}]bool

func NewSet() ISet {
	return &Set{}
}

func (s Set) Add(key interface{}) {
	s[key] = true
}

func (s Set) Remove(key interface{}) {
	delete(s, key)
}

func (s Set) Exists(key interface{}) bool {
	return s[key]
}

func (s Set) Size() int {
	return len(s)
}
