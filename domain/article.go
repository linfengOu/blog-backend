package domain

import (
	"gorm.io/gorm"
)

/*
 * models
 */
type ArticleEntity struct {
	gorm.Model
	Title   string     `gorm:"column:title"`
	Content string     `gorm:"column:content"`
	Author  UserEntity `gorm:"column:author;foreignKey:Email"`
}

/*
 * interface
 */
type ArticleRepository interface {
	FindOne(condition interface{}) (*ArticleEntity, error)
	Find(condition interface{}, page, pageSize int, columns ...string) ([]*ArticleEntity, error)

	Create(article *ArticleEntity) (*ArticleEntity, error)
	Update(data interface{}) (int64, error)

	DeleteOne(article *ArticleEntity) error
}

var articleRepository ArticleRepository

func GetArticleRepository() ArticleRepository {
	return articleRepository
}

func InitArticleRepository(ar ArticleRepository) {
	articleRepository = ar
}
