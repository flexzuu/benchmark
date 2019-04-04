package main

import (
	pb "github.com/flexzuu/thesis/example/grpc/post/post"
	"github.com/flexzuu/thesis/example/grpc/post/repo/entity"
)

func ToProto(p entity.Post) *pb.Post {
	return &pb.Post{
		ID:       p.ID,
		AuthorID: p.AuthorID,
		Content:  p.Content,
		Headline: p.Headline,
	}
}
