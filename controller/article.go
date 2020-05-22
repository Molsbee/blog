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
		context.JSON(err.StatusCode(), err.Details())
		return
	}

	if err := ac.articleService.Create(articleRequest); err != nil {
		context.JSON(err.StatusCode(), err.Details())
	}
}

func (ac *articleController) ListArticles(context *gin.Context) {
	articles, err := ac.articleService.List()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	abbreviatedArticleResponse := make([]model.AbbrArticleResponse, len(articles))
	for i, article := range articles {
		abbreviatedArticleResponse[i] = model.AbbrArticleResponse{
			ID:          article.ID,
			Title:       article.Title,
			Author:      article.Author,
			AbbrContent: article.Description(),
			CreatedDate: article.CreatedAt,
		}
	}

	context.JSON(200, abbreviatedArticleResponse)
}

func (ac *articleController) GetArticle(context *gin.Context) {
	articleID, errorResponse := parseArticleID(context)
	if errorResponse != nil {
		context.JSON(http.StatusBadRequest, errorResponse.Details())
		return
	}

	article, apiError := ac.articleService.Get(articleID)
	if apiError != nil {
		context.JSON(apiError.StatusCode(), apiError.Details())
		return
	}

	context.JSON(200, model.ArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		Author:      article.Author,
		Content:     article.Content,
		CreatedDate: article.CreatedAt,
	})
}

func (ac *articleController) UpdateArticle(context *gin.Context) {
	articleID, errorResponse := parseArticleID(context)
	if errorResponse != nil {
		context.JSON(http.StatusBadRequest, errorResponse.Details())
		return
	}

	articleRequest, err := parseArticleRequest(context)
	if err != nil {
		context.JSON(err.StatusCode(), err.Details())
		return
	}

	apiError := ac.articleService.Update(articleID, articleRequest)
	if apiError != nil {
		context.JSON(apiError.StatusCode(), apiError.Details())
	}
}

func parseArticleID(context *gin.Context) (articleID int, apiError model.ApplicationError) {
	articleID, conversionErr := strconv.Atoi(context.Param("articleID"))
	if conversionErr != nil {
		apiError = model.ErrorBuilder().
			StatusCode(http.StatusBadRequest).
			Message("invalid article id provided").
			Build()
	}
	return
}

func parseArticleRequest(context *gin.Context) (articleRequest model.ArticleRequest, err model.ApplicationError) {
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
