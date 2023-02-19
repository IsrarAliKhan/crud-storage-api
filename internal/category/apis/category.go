package api

import (
	"crud-storage-api/internal/category/models/orms"
	"crud-storage-api/internal/category/models/requests"
	"crud-storage-api/internal/category/models/responses"
	"crud-storage-api/internal/category/services"
	apis "crud-storage-api/shared/api"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Category struct {
	svc *services.CategoryService
}

func NewCategory(svc *services.CategoryService) *Category {
	return &Category{svc: svc}
}

func (api *Category) Routes(r chi.Router) {
	r.Route("/category", func(r chi.Router) {
		r.Post("/", api.AddCategory)
		r.Route("/{category_id}", func(r chi.Router) {
			r.Get("/", api.GetCategory)
			r.Patch("/", api.UpdateCategory)
			r.Delete("/", api.DeleteCategory)
		})
	})
}

// GetCategory is handler for route GET /category/{category_id}
func (api *Category) GetCategory(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "category_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `category_id` is required")
		return
	}

	// parameters type conversion
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call api service to fetch the category
	category, err := api.svc.GetCategory(uint64(categoryId))
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = responses.Category{
		Id:   category.Id,
		Name: category.Name,
		Type: category.Type,
	}

	// respond
	apis.RespondOk(w, res)
}

// AddCategory is handler for route POST /category
func (api *Category) AddCategory(w http.ResponseWriter, r *http.Request) {
	// parse and validate request
	var req requests.Category
	err := apis.ParseRequest(r, &req)
	if err != nil {
		apis.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// request conversion to storage model
	var category = orms.Category{
		Id:   req.Id,
		Name: req.Name,
		Type: req.Type,
	}

	// call api service to save the category
	categoryId, err := api.svc.SaveCategory(category)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      categoryId,
		Code:    http.StatusOK,
		Message: "category saved successfully",
	}

	// respond
	apis.RespondOk(w, res)
}

// UpdateCategory is handler for route PATCH /category/{category_id}
func (api *Category) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "category_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `category_id` is required")
		return
	}

	// parameters type conversion
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// parse and validate request
	var req requests.Category
	err = apis.ParseRequest(r, &req)
	if err != nil {
		apis.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// request conversion to storage model
	var category = orms.Category{
		Id:   uint64(categoryId),
		Name: req.Name,
		Type: req.Type,
	}

	// call api service to update the category
	err = api.svc.UpdateCategory(category)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      category.Id,
		Code:    http.StatusOK,
		Message: "category updated successfully",
	}

	// respond
	apis.RespondOk(w, res)
}

// DeleteCategory is handler for route DELETE /category/{category_id}
func (api *Category) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "category_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `category_id` is required")
		return
	}

	// parameters type conversion
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call api service to delete the category
	err = api.svc.DeleteCategory(uint64(categoryId))
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      uint64(categoryId),
		Code:    http.StatusOK,
		Message: "category deleted successfully",
	}

	// respond
	apis.RespondOk(w, res)
}
