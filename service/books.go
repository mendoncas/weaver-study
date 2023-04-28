package service

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ServiceWeaver/weaver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	client := getMongoClient()
	coll := client.Database("Books").Collection("books")
	res, err := coll.InsertOne(context.TODO(), book{Title: "Título", Author: "Autor", Description: "Descrição do livro"})
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("res.InsertedID: %v\n", res.InsertedID)
	}
	return err
}

// GET Books
func (b *books) FindBookByTitle(c context.Context, title string) error {
	client := getMongoClient()
	coll := client.Database("Books").Collection("books")
	var result bson.M
	err := coll.FindOne(c, bson.D{{"title", title}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	print(result)
	return err
}

func getMongoClient() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	fmt.Println("uri: ", uri)
	if uri == "" {
		// log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		fmt.Printf("erro")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}
