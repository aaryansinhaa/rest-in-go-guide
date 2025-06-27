package storage

type Storage interface {
	CreateBlock(name string, age int, address string, email string) (int64, error)
}
