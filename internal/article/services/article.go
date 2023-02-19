package services

import (
	"crud-storage-api/internal/article/models/orms"
	services "crud-storage-api/shared/service"
	"crud-storage-api/shared/storage"
	"fmt"
)

type ArticleService struct {
	strg storage.Medium
}

func NewArticleService(strg storage.Medium) *ArticleService {
	return &ArticleService{strg}
}

func (ArticleService) ServiceType() int { return services.ARTICLE_SERVICE }

func (s *ArticleService) GetArticle(id uint64) (orms.Article, error) {
	res, err := s.strg.Get(id)
	if err != nil {
		return orms.Article{}, err
	}
	article, ok := res.(orms.Article)
	if !ok {
		return orms.Article{}, fmt.Errorf("conversion failed")
	}
	return article, err
}

func (s *ArticleService) SaveArticle(m orms.Article) (uint64, error) {
	return s.strg.Save(m)
}

func (s *ArticleService) UpdateArticle(m orms.Article) error {
	return s.strg.Update(m)
}

func (s *ArticleService) DeleteArticle(id uint64) error {
	return s.strg.Delete(id)
}
