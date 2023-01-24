package set

type ISet interface {
	Add(key interface{})
	Remove(key interface{})
	Exists(key interface{}) bool
	Size() int
}
