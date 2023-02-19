package orms

import (
	storage_model "crud-storage-api/shared/storage/models"
	"fmt"
	"strconv"
	"strings"
)

type Category struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewCategory() *Category {
	return &Category{}
}

func (c Category) GetId() uint64 { return c.Id }
func (c Category) GetModel() storage_model.StorageORM {
	return Category{Id: c.Id, Name: c.Name, Type: c.Type}
}
func (Category) StorageType() int      { return storage_model.STORAGE_TYPE_FILE }
func (Category) GetEntityName() string { return "category" }

func (c Category) String() string {
	return fmt.Sprintf("%v, %v, %v\n", c.Id, c.Name, c.Type)
}

func (Category) Parse(s string) (storage_model.StorageORM, error) {
	if s == "" {
		return nil, fmt.Errorf("could not parse: %v", s)
	}

	s = strings.Trim(s, " \n")
	ss := strings.Split(s, ",")

	id, err := strconv.Atoi(ss[0])
	if err != nil {
		return Category{}, err
	}

	return Category{
		Id:   uint64(id),
		Name: ss[1],
		Type: ss[2],
	}, nil
}
