package services

const (
	ARTICLE_SERVICE = iota
	CATEGORY_SERVICE
	ITEM_SERVICE
)

type Service interface {
	ServiceType() int
}
