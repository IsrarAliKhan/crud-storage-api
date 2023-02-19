package main

import (
	"crud-storage-api/config"
	aApi "crud-storage-api/internal/article/apis"
	aSvc "crud-storage-api/internal/article/services"
	cApi "crud-storage-api/internal/category/apis"
	cOrm "crud-storage-api/internal/category/models/orms"
	cSvc "crud-storage-api/internal/category/services"
	iApi "crud-storage-api/internal/item/apis"
	"crud-storage-api/internal/item/db"
	iOrm "crud-storage-api/internal/item/models/orms"
	iSvc "crud-storage-api/internal/item/services"
	"crud-storage-api/internal/server"
	"crud-storage-api/shared/storage"
	storage_model "crud-storage-api/shared/storage/models"
	"log"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// storage
	m := make(map[uint64]storage_model.StorageORM)
	dynamic := storage.NewDynamicStorage(m)
	file := storage.NewFileStorage(config.FileName, cOrm.NewCategory())
	db := storage.NewDB(db.Conn(), iOrm.NewItem())

	// services
	articleSvc := aSvc.NewArticleService(dynamic)
	categorySvc := cSvc.NewCategoryService(file)
	itemSvc := iSvc.NewItemService(db)

	// apis
	articleApi := aApi.NewArticle(articleSvc)
	categoryApi := cApi.NewCategory(categorySvc)
	itemApi := iApi.NewItem(itemSvc)

	// server
	port, _ := strconv.Atoi(config.HttpPort)
	err := server.New(port, itemApi, categoryApi, articleApi).Start()
	if err != nil {
		log.Fatalf("start server err: %v", err)
	}
}
