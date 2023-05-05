package service

import (
	"context"
	"encoding/json"

	"github.com/ServiceWeaver/weaver"
	"github.com/mendoncas/weaver-study/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Books interface {
	RegisterBook(context.Context, Book) error
	FindBookByTitle(context.Context, string) ([]byte, error)
}

type books struct {
	weaver.Implements[Books]
	details      Details
	reviews      Reviews
	dbCollection *mongo.Collection
}

type Book struct {
	weaver.AutoMarshal
	Id          string `bson:"id"`
	Title       string `bson:"title"`
	Author      string `bson:"author"`
	Description string `bson:"description"`
}

func (b *books) Init(c context.Context) error {
	var err error
	b.details, err = weaver.Get[Details](b)
	b.reviews, err = weaver.Get[Reviews](b)
	b.dbCollection = client.GetMongoClient().Database("Books").Collection("books")
	return err
}

// POST Books
func (b *books) RegisterBook(c context.Context, data Book) error {
	_, err := b.dbCollection.InsertOne(context.TODO(), data)
	if err != nil {
		b.Logger().Error("erro ao inserir livro!", err)
	}
	return err
}

// GET Books
func (b *books) FindBookByTitle(c context.Context, title string) ([]byte, error) {
	var res Book
	err := b.dbCollection.FindOne(c, bson.D{{"title", title}}).Decode(&res)
	if err != nil {
		b.Logger().Error("erro ao buscar livro!", err)
	}
	r, _ := json.Marshal(res)

	return r, err
}
