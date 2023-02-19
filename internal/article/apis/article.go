package api

import (
	"crud-storage-api/internal/article/models/orms"
	"crud-storage-api/internal/article/models/requests"
	"crud-storage-api/internal/article/models/responses"
	"crud-storage-api/internal/article/services"
	apis "crud-storage-api/shared/api"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Article struct {
	svc *services.ArticleService
}

func NewArticle(svc *services.ArticleService) *Article {
	return &Article{svc: svc}
}

func (api *Article) Routes(r chi.Router) {
	r.Route("/article", func(r chi.Router) {
		r.Post("/", api.AddArticle)
		r.Route("/{article_id}", func(r chi.Router) {
			r.Get("/", api.GetArticle)
			r.Patch("/", api.UpdateArticle)
			r.Delete("/", api.DeleteArticle)
		})
	})
}

// GetArticle is handler for route GET /article/{article_id}
func (api *Article) GetArticle(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "article_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `article_id` is required")
		return
	}

	// parameters type conversion
	articleId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call api service to fetch the article
	article, err := api.svc.GetArticle(uint64(articleId))
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = responses.Article{
		Id:       article.Id,
		Name:     article.Name,
		Quantity: article.Quantity,
	}

	// respond
	apis.RespondOk(w, res)
}

// AddArticle is handler for route POST /article
func (api *Article) AddArticle(w http.ResponseWriter, r *http.Request) {
	// parse and validate request
	var req requests.Article
	err := apis.ParseRequest(r, &req)
	if err != nil {
		apis.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// request conversion to storage model
	var article = orms.Article{
		Id:       req.Id,
		Name:     req.Name,
		Quantity: req.Quantity,
	}

	// call api service to save the article
	articleId, err := api.svc.SaveArticle(article)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      articleId,
		Code:    http.StatusOK,
		Message: "article saved successfully",
	}

	// respond
	apis.RespondOk(w, res)
}

// UpdateArticle is handler for route PATCH /article/{article_id}
func (api *Article) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "article_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `article_id` is required")
		return
	}

	// parameters type conversion
	articleId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// parse and validate request
	var req requests.Article
	err = apis.ParseRequest(r, &req)
	if err != nil {
		apis.RespondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// request conversion to storage model
	var article = orms.Article{
		Id:       uint64(articleId),
		Name:     req.Name,
		Quantity: req.Quantity,
	}

	// call api service to update the article
	err = api.svc.UpdateArticle(article)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      article.Id,
		Code:    http.StatusOK,
		Message: "article updated successfully",
	}

	// respond
	apis.RespondOk(w, res)
}

// DeleteArticle is handler for route DELETE /article/{article_id}
func (api *Article) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	id := chi.URLParam(r, "article_id")
	if id == "" {
		apis.RespondErrorMessage(w, http.StatusBadRequest, "url parameter `article_id` is required")
		return
	}

	// parameters type conversion
	articleId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// call api service to delete the article
	err = api.svc.DeleteArticle(uint64(articleId))
	if err != nil {
		log.Println(err)
		apis.RespondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	var res = apis.DefaultResponse{
		Id:      uint64(articleId),
		Code:    http.StatusOK,
		Message: "article deleted successfully",
	}

	// respond
	apis.RespondOk(w, res)
}
