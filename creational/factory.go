package creational

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type Disk struct{}

func (s Disk) Open(fileName string) (io.ReadWriteCloser, error) {
	return nil, nil
}

type Memory struct{}

func (m Memory) Open(fileName string) (io.ReadWriteCloser, error) {
	return nil, nil
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case DiskStorage:
		return Disk{}
	case MemoryStorage:
		return Memory{}
	default:
		return nil
	}
}
