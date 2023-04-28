package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type Reviews interface {
	GetAllBookReviews(context.Context, string) error
}

type reviews struct {
	weaver.Implements[Reviews]
}

func (r *reviews) GetAllBookReviews(c context.Context, bookId string) error {
	fmt.Print(bookId)
	return errors.New("erro novo!")
}

// getRatings
// bookReviewsByID
