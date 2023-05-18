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
	// if ENV['ENABLE_EXTERNAL_BOOK_SERVICE'] === 'true' then
	//   # the ISBN of one of Comedy of Errors on the Amazon
	//   # that has Shakespeare as the single author
	//     isbn = '0486424618'
	//     return fetch_details_from_external_service(isbn, id, headers)
	// end
	return detail{Id: bookId, Author: "Shakespeare", Year: 487, Type: "paperback", Pages: 455, Publisher: "PublisherA", Language: "English"}, nil
}

func (d *details) AddBookDetails(c context.Context, s string) error {
	return errors.New("erro novo")
}
