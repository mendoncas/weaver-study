package service

import (
	"context"
	"errors"

	"github.com/ServiceWeaver/weaver"
)

type Details interface {
	AddBookDetails(context.Context, string) error
	FindDetailsByBookId(context.Context, string) error
}

type details struct {
	weaver.Implements[Details]
}

type detail struct {
	Date      string `bson:"date"`
	Publisher string `bson:"publisher"`
	BookID    string `bson:"book_id"`
}

func (d *details) FindDetailsByBookId(c context.Context, bookId string) error {
	return errors.New("erro novo")
}

func (d *details) AddBookDetails(c context.Context, s string) error {
	return errors.New("erro novo")
}

// GET ISBN
// GET BOOK DETAILS
// GET DETAILS FROM EXTERNAL SERVICE
