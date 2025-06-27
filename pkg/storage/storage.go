package storage

import "github.com/aaryansinhaa/miyuki/pkg/types"

type Storage interface {
	CreateBlockInStorage(m types.Miyuki) (int64, error)
	GetBlockByID(id int64) (*types.Miyuki, error)
}
