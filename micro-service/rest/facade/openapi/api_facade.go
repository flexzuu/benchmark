/*
 * Facade Service
 *
 * a facade service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthorDetail - Author Detail
func AuthorDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ListPosts - List Posts
func ListPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostDetail - Post Detail
func PostDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}