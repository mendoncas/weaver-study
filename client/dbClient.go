package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ServiceWeaver/weaver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbClient interface {
	StartConnection(context.Context) error
	Insert(context.Context, string, string, string) (string, error)
}

type dbClient struct {
	weaver.Implements[DbClient]
	mongo *mongo.Client
}

type StringInvertion struct {
	ReversedStr string `bson:"reversed_str"`
}

func (d *dbClient) Init(c context.Context) error {
	var err error
	d.mongo, err = d.getClient(c)
	return err
}

func (c *dbClient) getClient(_ context.Context) (*mongo.Client, error) {
	log := c.Logger()
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Info("NÃO CONSEGUI OBTER A URI DO BANCO!!")
	}
	log.Info(fmt.Sprintf("tentando conexão no endereço %s", uri))
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Error(fmt.Sprintf("Falha ao obter conexão! uri: %s", uri), err)
	}
	return client, err
}

func (c *dbClient) Insert(context context.Context,
	database string, collection string, content string) (string, error) {
	var err error
	client := c.mongo
	log := c.Logger()
	object := StringInvertion{ReversedStr: content}
	res, err := client.Database(database).Collection(collection).InsertOne(context, object)
	if err != nil {
		log.Error("erro ao inserir no banco!", err)
	} else {
		log.Info("res.InsertedID: %v\n", res.InsertedID)
	}
	return content, err
}

func (c *dbClient) StartConnection(_ context.Context) error {
	log := c.Logger()
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Info("falha ao buscar uri do mongo!!")
	}
	log.Info("tentando aqui conectar com a url: " + uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Error("falha ao conectar com o banco!", err)
	} else {
		log.Info("consegui conectar com o banco!")
	}

	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
	}

	if err != nil {
		return (err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return (err)
	}
	fmt.Printf("%s\n", jsonData)
	return err
}
