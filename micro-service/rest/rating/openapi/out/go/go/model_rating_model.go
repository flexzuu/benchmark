/*
 * Rating Service
 *
 * a rating service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RatingModel - A Rating
type RatingModel struct {

	Id int64 `json:"id,omitempty"`

	PostId int64 `json:"postId,omitempty"`

	Value int32 `json:"value,omitempty"`
}