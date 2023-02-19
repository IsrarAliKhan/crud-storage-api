package storage

import (
	"fmt"
	storage_model "crud-storage-api/shared/storage/models"
)

var strg map[uint64]storage_model.StorageORM

type DynamicStorage struct{}

func NewDynamicStorage(x map[uint64]storage_model.StorageORM) *DynamicStorage {
	strg = x
	return &DynamicStorage{}
}

func (DynamicStorage) Get(id uint64) (storage_model.StorageORM, error) {
	if _, ok := strg[id]; ok {
		return strg[id], nil
	} else {
		return nil, fmt.Errorf("reocrd not found for %v", id)
	}
}

func (DynamicStorage) Save(m storage_model.StorageORM) (uint64, error) {
	if _, ok := strg[m.GetId()]; !ok {
		strg[m.GetId()] = m.GetModel()
		return m.GetId(), nil
	}
	return 0, fmt.Errorf("record already exists")
}

func (DynamicStorage) Update(m storage_model.StorageORM) error {
	if _, ok := strg[m.GetId()]; ok {
		strg[m.GetId()] = m.GetModel()
		return nil
	}
	return fmt.Errorf("reocrd not found for %v", m.GetId())
}

func (DynamicStorage) Delete(id uint64) error {
	delete(strg, id)
	return nil
}
