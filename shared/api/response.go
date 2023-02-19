package api

type DefaultResponse struct {
	Id      uint64 `json:"id"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
