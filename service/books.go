package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/mendoncas/weaver-study/client"
	"go.mongodb.org/mongo-driver/bson"
)

type Books interface {
	RegisterBook(context.Context) error
	FindBookByTitle(context.Context, string) error
}

type books struct {
	weaver.Implements[Books]
	details Details
	reviews Reviews
}

type book struct {
	Title       string `bson:"title"`
	Author      string `bson:"author"`
	Description string `bson:"description"`
}

func (b *books) Init(c context.Context) error {
	b.details, _ = weaver.Get[Details](b)
	b.reviews, _ = weaver.Get[Reviews](b)
	return errors.New("erro!")
}

// POST Books
func (b *books) RegisterBook(c context.Context) error {
	fmt.Printf("fui chamada!")
	log := b.Logger()
	log.Info("tentando inserir um livro aqui...")
	cl := client.GetMongoClient()
	coll := cl.Database("Books").Collection("books")
	res, err := coll.InsertOne(context.TODO(), book{Title: "Título", Author: "Autor", Description: "Descrição do livro"})
	if err != nil {
		log.Error("erro ao inserir livro!", err)
	} else {
		fmt.Printf("res.InsertedID: %v\n", res.InsertedID)
	}
	return err
}

// GET Books
func (b *books) FindBookByTitle(c context.Context, title string) error {
	cl := client.GetMongoClient()

	coll := cl.Database("Books").Collection("books")
	var result bson.M
	err := coll.FindOne(c, bson.D{{"title", title}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	print(result)
	return err
}
