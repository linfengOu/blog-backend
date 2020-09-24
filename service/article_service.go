package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github/linfengOu/write-backend/domain"
	"gorm.io/gorm"
	"time"
)

/*
 * models
 */
type ArticleAbstractModel struct {
	ID        uint
	Title     string
	Author    domain.UserAbstractModel
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleModel struct {
	ID        uint
	Title     string
	Content   string
	Author    domain.UserAbstractModel
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (model *ArticleModel) String() string {
	return fmt.Sprintf("ArticleModel[ID: %d, Title: %s, Author: %s]", model.ID, model.Title, model.Author.Username)
}

/*
 * interface
 */
type ArticleService interface {
	Get(id uint) (*ArticleModel, error)
	GetAll(page, pageSize int) ([]*ArticleAbstractModel, error)
	Create(article *ArticleModel) (*ArticleModel, error)
}

var articleService ArticleService

func GetArticleService() ArticleService {
	return articleService
}

func InitArticleService(as ArticleService) {
	articleService = as
}

/*
 * implementation
 */
type ArticleServiceImpl struct {
}

func (as *ArticleServiceImpl) Get(id uint) (*ArticleModel, error) {
	articleEntity, err := domain.GetArticleRepository().FindOne(map[string]interface{}{"ID": id})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find Article by ID %d failed", id))
	}
	return toModel(articleEntity), nil
}

func (as *ArticleServiceImpl) GetAll(page, pageSize int) ([]*ArticleAbstractModel, error) {
	articleEntities, err := domain.GetArticleRepository().Find(&ArticleAbstractModel{}, page, pageSize, "Content")
	result := make([]*ArticleAbstractModel, len(articleEntities))

	if err != nil {
		return nil, errors.Wrap(err, "find all failed")
	}

	for _, articleEntity := range articleEntities {
		result = append(result, toAbstractModel(articleEntity))
	}
	return result, nil
}

func (as *ArticleServiceImpl) Create(article *ArticleModel) (*ArticleModel, error) {
	articleEntity, err := domain.GetArticleRepository().Create(toEntity(article))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("create failed %s", article))
	}
	return toModel(articleEntity), nil
}

/*
 * mapper
 */
func toEntity(model *ArticleModel) *domain.ArticleEntity {
	return &domain.ArticleEntity{
		Model: gorm.Model{
			ID: model.ID,
		},
		Title:   model.Title,
		Content: model.Content,
		Author: domain.UserEntity{
			Model: gorm.Model{
				ID: model.Author.ID,
			},
		},
	}
}

func toModel(entity *domain.ArticleEntity) *ArticleModel {
	return &ArticleModel{
		ID:      entity.ID,
		Title:   entity.Title,
		Content: entity.Content,
		Author: domain.UserAbstractModel{
			ID:       entity.Author.ID,
			Username: entity.Author.Username,
		},
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func toAbstractModel(entity *domain.ArticleEntity) *ArticleAbstractModel {
	return &ArticleAbstractModel{
		ID:    entity.ID,
		Title: entity.Title,
		Author: domain.UserAbstractModel{
			ID:       entity.Author.ID,
			Username: entity.Author.Username,
		},
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
