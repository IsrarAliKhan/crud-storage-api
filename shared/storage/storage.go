package storage

import storage_model "crud-storage-api/shared/storage/models"

type Medium interface {
	Get(uint64) (storage_model.StorageORM, error)
	Save(storage_model.StorageORM) (uint64, error)
	Update(storage_model.StorageORM) error
	Delete(uint64) error
}
