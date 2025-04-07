package main

import (
	"myapp/internal/api/handler"
	"myapp/internal/api/router"
	"myapp/internal/config"
	"myapp/internal/domain"
	"myapp/internal/repository/mysql"
	"myapp/internal/service"
)

func main() {
	// Initialize config
	cfg := config.NewConfig()

	// Initialize database
	db := config.InitDB(cfg)

	// Auto Migrate the schema
	db.AutoMigrate(&domain.Post{})

	// Initialize repository
	postRepo := mysql.NewPostRepository(db)

	// Initialize service
	postService := service.NewPostService(postRepo)

	// Initialize handlers
	postHandler := handler.NewPostHandler(postService)
	healthHandler := handler.NewHealthHandler()

	// Setup router
	r := router.SetupRouter(postHandler, healthHandler)

	// Run the server
	r.Run(":8080")
}
