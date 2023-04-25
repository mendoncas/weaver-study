package service

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ServiceWeaver/weaver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Books interface {
	// RegisterBook(context.Context) (*mongo.InsertOneResult, error)
	Test(context.Context) error
}

type books struct {
	weaver.Implements[Books]
}

type book struct {
	title, author, description string
}

func Test(c context.Context) error {
	fmt.Println("fui chamada!")
	return errors.New("error!")
}

// POST Books
func RegisterBook(c context.Context) (*mongo.InsertOneResult, error) {
	client := getMongoClient()
	coll := client.Database("Books").Collection("books")
	res, err := coll.InsertOne(context.TODO(), book{title: "Título", author: "Autor", description: "Descrição do livro"})
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("res.InsertedID: %v\n", res.InsertedID)
	}
	return res, err
}

// GET Books

func getMongoClient() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		// log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		fmt.Printf("erro")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return client
}
