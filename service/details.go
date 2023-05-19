package service

import (
	"context"
	"errors"

	"github.com/ServiceWeaver/weaver"
)

type detail struct {
	// Date      string `bson:"date"`
	// Publisher string `bson:"publisher"`
	// BookID    string `bson:"book_id"`
	weaver.AutoMarshal
	Id        string
	Author    string
	Year      int
	Type      string
	Pages     int
	Publisher string
	Language  string
	ISBN      string
}

type Details interface {
	// AddBookDetails(context.Context, string) error
	FindDetailsByBookId(context.Context, string) (detail, error)
}

type details struct {
	weaver.Implements[Details]
}

func (d *details) FindDetailsByBookId(c context.Context, bookId string) (detail, error) {
	return detail{Id: bookId, Author: "Shakespeare", Year: 487, Type: "paperback", Pages: 455, Publisher: "PublisherA", Language: "English"}, nil
}

func (d *details) AddBookDetails(c context.Context, s string) error {
	return errors.New("erro novo")
}
