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
		context.JSON(http.StatusInternalServerError, model.ApiErrorResponse{
			Message: "unable to parse request",
		})
		return
	}
	ac.articleService.Create(articleRequest)
}

func (ac *articleController) ListArticles(context *gin.Context) {
	articles := ac.articleService.List()
	context.JSON(200, articles)
}

func (ac *articleController) GetArticle(context *gin.Context) {
	articleID, errorResponse := parseArticleID(context)
	if errorResponse != nil {
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	context.JSON(200, ac.articleService.Get(articleID))
}

func (ac *articleController) UpdateArticle(context *gin.Context) {
	articleID, errorResponse := parseArticleID(context)
	if errorResponse != nil {
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	articleRequest, err := parseArticleRequest(context)
	if err != nil {
		context.JSON(500, model.ApiErrorResponse{
			Message: "unable to parse request",
		})
		return
	}
	ac.articleService.Update(articleID, articleRequest)
}

func parseArticleID(context *gin.Context) (articleID int, errorResponse *model.ApiErrorResponse) {
	articleID, err := strconv.Atoi(context.Query("articleID"))
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
