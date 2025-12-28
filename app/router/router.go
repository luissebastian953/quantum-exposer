package router

import (
	"net/http"
	"quantum-exposer/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(postController controllers.PostController) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok", "service": "quantum-exposer"})
	})

	v1 := router.Group("/api/v1")
	{
		// GET /api/v1/posts?limit=20&page=1&tags=...
		v1.GET("/posts", postController.ListPosts)

		// GET /api/v1/posts/[10431619] (using Gin's path parameter ':id')
		v1.GET("/posts/:id", postController.GetPostByID)

		// GET /api/v1/tags?limit=20&page=1&name=...
		v1.GET("/tags", postController.GetTags)

		// GET /api/v1/tags/matches?limit=20&page=1&keyword=...
		v1.GET("/tags/matches", postController.GetTagsMatches)

		// GET /api/v1/tags/autocomplete?name=...
		v1.GET("/tags/autocomplete", postController.AutocompleteTags)

		// GET /api/v1/artists?limit=20&page=1&name=...
		v1.GET("/artists", postController.GetArtists)

		// Add other endpoints here... (e.g., tags)
	}

	return router
}
