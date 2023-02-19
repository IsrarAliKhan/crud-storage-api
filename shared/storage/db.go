package storage

import (
	storage_model "crud-storage-api/shared/storage/models"

	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
	orm  storage_model.StorageORM
}

func NewDB(conn *gorm.DB, orm storage_model.StorageORM) *DB {
	return &DB{conn, orm}
}

func (s DB) Get(id uint64) (storage_model.StorageORM, error) {
	var m = s.orm.GetModel()
	err := s.conn.
		Model(s.orm.GetModel()).
		Where("id = ?", id).
		First(&m).
		Error
	return m, err
}

func (s DB) Save(m storage_model.StorageORM) (uint64, error) {
	err := s.conn.Create(m.GetModel()).Error
	return m.GetId(), err
}

func (s DB) Update(m storage_model.StorageORM) error {
	if err := s.conn.
		Where("id = ?", m.GetId()).
		First(m.GetModel()).
		Error; err != nil {
		return err
	}
	return s.conn.Save(m.GetModel()).Error
}

func (s DB) Delete(id uint64) error {
	if err := s.conn.
		Where("id = ?", id).
		First(s.orm.GetModel()).
		Error; err != nil {
		return err
	}
	return s.conn.Where("id = ?", id).Delete(s.orm.GetModel()).Error
}
