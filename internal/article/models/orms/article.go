package orms

import (
	storage_model "crud-storage-api/shared/storage/models"
	"fmt"
	"strconv"
	"strings"
)

type Article struct {
	Id       uint64  `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}

func NewArticle() *Article {
	return &Article{}
}

func (a Article) GetId() uint64 { return a.Id }
func (a Article) GetModel() storage_model.StorageORM {
	return Article{Id: a.Id, Name: a.Name, Quantity: a.Quantity}
}
func (Article) StorageType() int      { return storage_model.STORAGE_TYPE_DYNAMIC }
func (Article) GetEntityName() string { return "article" }

func (c Article) String() string {
	return fmt.Sprintf("%v, %v, %v\n", c.Id, c.Name, c.Quantity)
}

func (Article) Parse(s string) (storage_model.StorageORM, error) {
	if s == "" {
		return nil, fmt.Errorf("could not parse: %v", s)
	}

	s = strings.Trim(s, " \n")
	ss := strings.Split(s, ",")

	id, err := strconv.Atoi(ss[0])
	if err != nil {
		return Article{}, err
	}

	quantity, err := strconv.ParseFloat(ss[2], 64)
	if err != nil {
		return Article{}, err
	}

	return Article{
		Id:       uint64(id),
		Name:     ss[1],
		Quantity: quantity,
	}, nil
}
