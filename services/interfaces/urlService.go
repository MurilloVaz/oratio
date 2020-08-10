package interfaces

import (
	e "github.com/MurilloVaz/oratio/entities"
)

type IUrlService interface {
	InsertOne(url *e.Url) error
	FindOne(id *string) (*e.Url, error)
	GetAll() ([]*e.Url, error)
}
