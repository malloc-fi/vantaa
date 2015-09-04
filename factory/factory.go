package factory

import (
	"github.com/nathandao/vantaa/model"
)

type Factory struct{}

func NewFactory() Factory {
	return Factory{}
}

func (f *Factory) Users() []*model.User {
	users := []*model.User{
		&model.User{
			Name:     "admin",
			Email:    "admin@example.com",
			Password: "adminpass",
		},
		&model.User{
			Name:     "johndoe",
			Email:    "jhondoe@example.com",
			Password: "johndoepass",
		},
		&model.User{
			Name:     "foo",
			Email:    "foo@example.com",
			Password: "foopass",
		},
		&model.User{
			Name:     "bar",
			Email:    "bar@example.com",
			Password: "barpass",
		},
	}
	return users
}
