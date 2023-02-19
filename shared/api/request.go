package api

type Request interface {
	Validate() error
}
