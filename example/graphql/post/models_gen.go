// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package post

type PostCreateInput struct {
	AuthorID string `json:"authorId"`
	Headline string `json:"headline"`
	Content  string `json:"content"`
}

type PostDeleteInput struct {
	ID string `json:"id"`
}

type PostDeletePayload struct {
	DeletedPostID string `json:"deletedPostId"`
}
