package controller

import (
	"github.com/Molsbee/blog/model"
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/gin"
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
		context.JSON(http.StatusBadRequest, model.ApiErrorResponse{
			Message: "unable to parse request",
		})
		return
	}

	if err := ac.articleService.Create(articleRequest); err != nil {
		context.JSON(http.StatusInternalServerError, model.ApiErrorResponse{
			Message: err.Error(),
		})
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

	article, err := ac.articleService.Get(articleID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.ApiErrorResponse{
			Message: err.Error(),
		})
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
		context.JSON(http.StatusBadRequest, model.ApiErrorResponse{
			Message: "unable to parse request",
		})
		return
	}
	ac.articleService.Update(articleID, articleRequest)
}

func parseArticleID(context *gin.Context) (articleID int, errorResponse *model.ApiErrorResponse) {
	articleID, err := strconv.Atoi(context.Param("articleID"))
	if err != nil {
		errorResponse = &model.ApiErrorResponse{
			Message: "invalid article id provided",
		}
	}
	return
}

func parseArticleRequest(context *gin.Context) (articleRequest model.ArticleRequest, err error) {
	err = context.BindJSON(&articleRequest)
	return
}
