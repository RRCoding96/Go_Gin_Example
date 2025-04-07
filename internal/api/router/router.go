package router

import (
	"myapp/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(postHandler *handler.PostHandler, healthHandler *handler.HealthHandler) *gin.Engine {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", healthHandler.Check)

	// Posts routes
	posts := r.Group("/posts")
	{
		posts.GET("", postHandler.GetPosts)
		posts.GET("/:id", postHandler.GetPost)
		posts.POST("", postHandler.CreatePost)
		posts.PUT("/:id", postHandler.UpdatePost)
		posts.DELETE("/:id", postHandler.DeletePost)
	}

	return r
}
