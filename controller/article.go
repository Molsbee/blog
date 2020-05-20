package controller

import (
	"github.com/Molsbee/blog/model"
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type articleController struct {
	articleService *service.ArticleService
}

func NewArticleController(articleService *service.ArticleService) *articleController {
	return &articleController{articleService: articleService}
}

func (ac *articleController) Create(context *gin.Context) {
	articleRequest, err := parseArticleRequest(context)
	if err != nil {
		context.JSON(err.StatusCode(), err)
		return
	}

	if err := ac.articleService.Create(articleRequest); err != nil {
		context.JSON(err.StatusCode(), err)
	}
}

func (ac *articleController) ListArticles(context *gin.Context) {
	articles, err := ac.articleService.List()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	context.JSON(200, articles)
}

func (ac *articleController) GetArticle(context *gin.Context) {
	articleID, errorResponse := parseArticleID(context)
	if errorResponse != nil {
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	article, apiError := ac.articleService.Get(articleID)
	if apiError != nil {
		context.JSON(apiError.StatusCode(), apiError)
		return
	}

	context.JSON(200, article)
}

func (ac *articleController) UpdateArticle(context *gin.Context) {
	articleID, errorResponse := parseArticleID(context)
	if errorResponse != nil {
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	articleRequest, err := parseArticleRequest(context)
	if err != nil {
		context.JSON(err.StatusCode(), err)
		return
	}

	apiError := ac.articleService.Update(articleID, articleRequest)
	if apiError != nil {
		context.JSON(apiError.StatusCode(), apiError)
	}
}

func parseArticleID(context *gin.Context) (articleID int, apiError model.ApiError) {
	articleID, conversionErr := strconv.Atoi(context.Param("articleID"))
	if conversionErr != nil {
		apiError = model.ErrorBuilder().
			StatusCode(http.StatusBadRequest).
			Message("invalid article id provided").
			Build()
	}
	return
}

func parseArticleRequest(context *gin.Context) (articleRequest model.ArticleRequest, err model.ApiError) {
	bindError := context.BindJSON(&articleRequest)
	if bindError != nil {
		log.Printf("failed to bind json to article request - %s\n", bindError)
		err = model.ErrorBuilder().
			StatusCode(http.StatusBadRequest).
			Message("unable to parse request body").
			Build()
	}

	return
}
