package controllers

import (
	"net/http"
	"quantum-exposer/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	Repo usecase.PostRepository
}

func NewPostController(repo usecase.PostRepository) *PostController {
	return &PostController{
		Repo: repo,
	}
}

func (c *PostController) ListPosts(ctx *gin.Context) {
	tags := ctx.QueryArray("tags")
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "20")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	posts, err := c.Repo.FetchPosts(usecase.SearchCriteriaPost{
		Tags:  tags,
		Page:  page,
		Limit: limit,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (c *PostController) GetPostByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})

		return
	}

	post, err := c.Repo.FetchPostByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": post})
}

func (c *PostController) GetTags(ctx *gin.Context) {
	names := ctx.QueryArray("names")
	limitStr := ctx.DefaultQuery("limit", "200")
	pageStr := ctx.DefaultQuery("page", "1")
	order := ctx.DefaultQuery("order", "post_count")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	tags, err := c.Repo.FetchTagsByName(usecase.SearchCriteriaTag{
		Names: names,
		Limit: limit,
		Page:  page,
		Order: order,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tags": tags})
}

func (c *PostController) GetTagsMatches(ctx *gin.Context) {
	keyword := ctx.DefaultQuery("keyword", "")
	limitStr := ctx.DefaultQuery("limit", "200")
	pageStr := ctx.DefaultQuery("page", "1")
	order := ctx.DefaultQuery("order", "count")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	if keyword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Keyword parameter is required"})
		return
	}

	tags, err := c.Repo.FetchTagsByNameMatches(usecase.SearchCriteriaTag{
		NamePrefix: keyword,
		Limit:      limit,
		Page:       page,
		Order:      order,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tags": tags})
}

func (c *PostController) GetArtists(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")
	limitStr := ctx.DefaultQuery("limit", "50")
	pageStr := ctx.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	artists, err := c.Repo.FetchArtists(usecase.SearchCriteriaArtist{
		Name:  name,
		Limit: limit,
		Page:  page,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"artists": artists})
}

func (c *PostController) GetRandomPost(ctx *gin.Context) {
	// Implementation for fetching a random post can be added here
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (c *PostController) AutocompleteTags(ctx *gin.Context) {
	// Implementation for tag autocomplete can be added here
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}
