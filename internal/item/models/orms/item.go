package orms

import (
	storage_model "crud-storage-api/shared/storage/models"
	"fmt"
	"strconv"
	"strings"
)

type Item struct {
	Id    uint64 `json:"id" gorm:"id"`
	Name  string `json:"name" gorm:"name"`
	Price uint64 `json:"price" gorm:"price"`
}

func NewItem() *Item {
	return &Item{}
}

func (i Item) GetId() uint64                      { return i.Id }
func (i Item) GetModel() storage_model.StorageORM { return &Item{i.Id, i.Name, i.Price} }
func (Item) StorageType() int                     { return storage_model.STORAGE_TYPE_DATABASE }
func (Item) GetEntityName() string                { return "item" }

func (i Item) String() string {
	return fmt.Sprintf("%v, %v, %v\n", i.Id, i.Name, i.Price)
}

func (Item) Parse(s string) (storage_model.StorageORM, error) {
	if s == "" {
		return nil, fmt.Errorf("could not parse: %v", s)
	}

	s = strings.Trim(s, " \n")
	ss := strings.Split(s, ",")

	id, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil, err
	}

	price, err := strconv.Atoi(ss[2])
	if err != nil {
		return nil, err
	}

	return &Item{
		Id:    uint64(id),
		Name:  ss[1],
		Price: uint64(price),
	}, nil
}
