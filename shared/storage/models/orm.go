package models

const (
	STORAGE_TYPE_FILE = iota
	STORAGE_TYPE_DYNAMIC
	STORAGE_TYPE_DATABASE
)

type StorageORM interface {
	GetId() uint64
	GetModel() StorageORM
	StorageType() int
	GetEntityName() string
	String() string
	Parse(s string) (StorageORM, error)
}
