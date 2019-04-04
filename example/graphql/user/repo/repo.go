package repo

import "github.com/flexzuu/thesis/example/graphql/user/repo/entity"

type User interface {
	Get(ID int) (entity.User, error)
	Create(Email string, Name string) (entity.User, error)
	Delete(ID int) error
}
