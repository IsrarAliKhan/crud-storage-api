package services

import (
	"crud-storage-api/internal/item/models/orms"
	services "crud-storage-api/shared/service"
	"crud-storage-api/shared/storage"
	"fmt"
)

type ItemService struct {
	strg storage.Medium
}

func NewItemService(strg storage.Medium) *ItemService {
	return &ItemService{strg}
}

func (ItemService) ServiceType() int { return services.ITEM_SERVICE }

func (s *ItemService) GetItem(id uint64) (*orms.Item, error) {
	res, err := s.strg.Get(id)
	item, ok := res.(*orms.Item)
	if !ok {
		return nil, fmt.Errorf("conversion failed")
	}
	return item, err
}

func (s *ItemService) SaveItem(m orms.Item) (uint64, error) {
	return s.strg.Save(m)
}

func (s *ItemService) UpdateItem(m orms.Item) error {
	return s.strg.Update(m)
}

func (s *ItemService) DeleteItem(id uint64) error {
	return s.strg.Delete(id)
}
