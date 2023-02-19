package api

import (
	"crud-storage-api/internal/item/models/orms"
	"crud-storage-api/internal/item/models/requests"
	"crud-storage-api/internal/item/models/responses"
	"crud-storage-api/internal/item/services"
	apis "crud-storage-api/shared/api"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Item struct {
	svc *services.ItemService
}

func NewItem(svc *services.ItemService) *Item {
	return &Item{svc: svc}
}

func (api *Item) Routes(r chi.Router) {
	r.Route("/item", func(r chi.Router) {
		r.Post("/", api.AddItem)
		r.Route("/{item_id}", func(r chi.Router) {
			r.Get("/", api.GetItem)
			r.Patch("/", api.UpdateItem)
			r.Delete("/", api.DeleteItem)
		})
	})
}

// GetItem is handler for route GET /item/{item_id}
func (api *Item) GetItem(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "item_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `item_id` is required")
		return
	}

	// parameters type conversion
	itemId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call api service to fetch the item
	item, err := api.svc.GetItem(uint64(itemId))
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = responses.Item{
		Id:    item.Id,
		Name:  item.Name,
		Price: item.Price,
	}

	// respond
	apis.RespondOk(w, res)
}

// AddItem is handler for route POST /item
func (api *Item) AddItem(w http.ResponseWriter, r *http.Request) {
	// parse and validate request
	var req requests.Item
	err := apis.ParseRequest(r, &req)
	if err != nil {
		apis.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// request conversion to storage model
	var item = orms.Item{
		Id:    req.Id,
		Name:  req.Name,
		Price: req.Price,
	}

	// call api service to save the item
	itemId, err := api.svc.SaveItem(item)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      itemId,
		Code:    http.StatusOK,
		Message: "item saved successfully",
	}

	// respond
	apis.RespondOk(w, res)
}

// UpdateItem is handler for route PATCH /item/{item_id}
func (api *Item) UpdateItem(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "item_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `item_id` is required")
		return
	}

	// parameters type conversion
	itemId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// parse and validate request
	var req requests.Item
	err = apis.ParseRequest(r, &req)
	if err != nil {
		apis.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// request conversion to storage model
	var item = orms.Item{
		Id:    uint64(itemId),
		Name:  req.Name,
		Price: req.Price,
	}

	// call api service to update the item
	err = api.svc.UpdateItem(item)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      item.Id,
		Code:    http.StatusOK,
		Message: "item updated successfully",
	}

	// respond
	apis.RespondOk(w, res)
}

// DeleteItem is handler for route DELETE /item/{item_id}
func (api *Item) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "item_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `item_id` is required")
		return
	}

	// parameters type conversion
	itemId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call api service to delete the item
	err = api.svc.DeleteItem(uint64(itemId))
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      uint64(itemId),
		Code:    http.StatusOK,
		Message: "item deleted successfully",
	}

	// respond
	apis.RespondOk(w, res)
}
