package services

import (
	"crud-storage-api/internal/category/models/orms"
	services "crud-storage-api/shared/service"
	"crud-storage-api/shared/storage"
	"fmt"
	"log"
)

type CategoryService struct {
	strg storage.Medium
}

func NewCategoryService(strg storage.Medium) *CategoryService {
	return &CategoryService{strg}
}

func (CategoryService) ServiceType() int { return services.CATEGORY_SERVICE }

func (s *CategoryService) GetCategory(id uint64) (orms.Category, error) {
	res, err := s.strg.Get(id)
	log.Println(err)
	if err != nil {
		return orms.Category{}, err
	}
	category, ok := res.(orms.Category)
	if !ok {
		return orms.Category{}, fmt.Errorf("conversion failed")
	}
	return category, err
}

func (s *CategoryService) SaveCategory(m orms.Category) (uint64, error) {
	return s.strg.Save(m)
}

func (s *CategoryService) UpdateCategory(m orms.Category) error {
	return s.strg.Update(m)
}

func (s *CategoryService) DeleteCategory(id uint64) error {
	return s.strg.Delete(id)
}
