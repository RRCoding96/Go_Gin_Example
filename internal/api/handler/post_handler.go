package handler

import (
	"myapp/internal/domain"
	"myapp/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

// GetPosts returns all posts
func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost returns a post by ID
func (h *PostHandler) GetPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID 형식입니다."})
		return
	}

	post, err := h.postService.GetPost(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "게시글을 찾을 수 없습니다."})
		return
	}
	c.JSON(http.StatusOK, post)
}

// CreatePost creates a new post
func (h *PostHandler) CreatePost(c *gin.Context) {
	var post domain.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPost, err := h.postService.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdPost)
}

// UpdatePost updates an existing post
func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID 형식입니다."})
		return
	}

	var post domain.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.ID = uint(id)
	updatedPost, err := h.postService.UpdatePost(post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "게시글을 찾을 수 없습니다."})
		return
	}
	c.JSON(http.StatusOK, updatedPost)
}

// DeletePost deletes a post
func (h *PostHandler) DeletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID 형식입니다."})
		return
	}

	if err := h.postService.DeletePost(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "게시글을 찾을 수 없습니다."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "게시글이 삭제되었습니다."})
}
