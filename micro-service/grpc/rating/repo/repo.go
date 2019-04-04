package repo

import "github.com/flexzuu/thesis/micro-service/grpc/rating/repo/entity"

type Rating interface {
	GetById(ID int64) (entity.Rating, error)
	ListOfPost(PostID int64) ([]entity.Rating, error)
	Create(PostID int64, value int32) (entity.Rating, error)
	Delete(ID int64) error
}
