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
	post "github.com/flexzuu/thesis/example/openapi/post/openapi"
)

// PostListModel - a list of posts
type PostListModel struct {
	Posts []post.PostModel `json:"posts"`
}