package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github/linfengOu/write-backend/service"
	"net/http"
	"strconv"
)

const ArticleKey = "articles"

func ArticleRetrieve(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	articleModel, err := service.GetArticleService().Get(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, articleModel)
}

func ArticleList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, err2 := strconv.Atoi(c.DefaultQuery("size", "20"))
	if err != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.New("invalid paging params"))
	}
	abstractModels, err := service.GetArticleService().GetAll(page, size)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, abstractModels)
}

func ArticleCreate(c *gin.Context) {

}
