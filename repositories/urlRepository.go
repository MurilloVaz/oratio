package repositories

import (
	"context"

	data "github.com/MurilloVaz/oratio/data"
	e "github.com/MurilloVaz/oratio/entities"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	COLLECTION = "urls"
)

type UrlRepository struct {
	ctx *data.DataContext
}

func (u *UrlRepository) GetAll() ([]*e.Url, error) {
	var urls []*e.Url

	cur, err := u.ctx.Collection(COLLECTION).Find(context.TODO(), bson.D{{}})

	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var elem e.Url
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		urls = append(urls, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return urls, err
}

func (u *UrlRepository) FindOne(id *string) (*e.Url, error) {
	var url e.Url

	err := u.ctx.Collection(COLLECTION).FindOne(context.TODO(), bson.M{"_id": id}).Decode(&url)

	if err != nil {
		return nil, err
	}
	return &url, err
}

func (u *UrlRepository) InsertOne(url *e.Url) error {
	_, err := u.ctx.Collection(COLLECTION).InsertOne(context.TODO(), url)
	if err != nil {
		return err
	}
	return nil
}

func NewUrlRepository(ctx *data.DataContext) *UrlRepository {
	return &UrlRepository{ctx}
}

func (u *UrlRepository) Disconnect() error {
	return u.ctx.Disconnect()
}
