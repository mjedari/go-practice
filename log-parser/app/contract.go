package app

type IStorage interface {
	Persist([]byte) error
}
