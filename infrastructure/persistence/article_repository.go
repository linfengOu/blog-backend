package persistence

import (
	"github/linfengOu/write-backend/domain"
)

type ArticleRepository struct {
}

func (r *ArticleRepository) FindOne(condition interface{}) (*domain.ArticleEntity, error) {
	// condition can be domain.ArticleEntity struct
	db := GetDB()
	var article *domain.ArticleEntity
	err := db.Where(condition).First(article).Error
	return article, err
}

func (r *ArticleRepository) Find(condition interface{}, page, pageSize int, columns ...string) ([]*domain.ArticleEntity, error) {
	db := GetDB()
	if page < 1 {
		page = 1
	}

	switch {
	case pageSize == 0:
		pageSize = 30
	case pageSize > 500:
		pageSize = 500
	}

	var articles []*domain.ArticleEntity
	err := db.Where(condition).Offset((page - 1) * pageSize).Limit(pageSize).Find(articles).Omit(columns...).Error

	return articles, err
}

func (r *ArticleRepository) Create(article *domain.ArticleEntity) (*domain.ArticleEntity, error) {
	db := GetDB()
	err := db.Create(article).Error
	return article, err
}

func (r *ArticleRepository) Update(data interface{}) (int64, error) {
	// data can be struct(only update no-nil field) or map
	db := GetDB()
	tx := db.Model(&domain.ArticleEntity{}).Updates(data)
	return tx.RowsAffected, tx.Error
}

func (r *ArticleRepository) DeleteOne(article *domain.ArticleEntity) error {
	db := GetDB()
	err := db.Delete(article).Error
	return err
}
