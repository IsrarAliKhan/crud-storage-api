package api

import (
	"github.com/go-chi/chi"
)

type API interface {
	Routes(r chi.Router)
}
