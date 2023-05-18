package service

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type Ratings interface {
	PutLocalReviews(context.Context, string, string) error
	GetRatingsByBookId(context.Context, string) (rating, error)
}

type rating struct {
	weaver.AutoMarshal
	Id     string
	Rating []string
}

type ratings struct {
	weaver.Implements[Ratings]
	// userAddedRatings []rating
	userAddedRatings map[string][]string
}

func (r *ratings) Init(c context.Context) error {
	var err error
	// r.userAddedRatings = []rating{}
	r.userAddedRatings = map[string][]string{}

	return err
}

func (r *ratings) GetRatingsByBookId(c context.Context, bookId string) (rating, error) {
	r.printUserAddedRatings()
	bookRatings := r.userAddedRatings[bookId]
	fmt.Printf(bookRatings[0])
	return rating{Id: bookId, Rating: bookRatings}, nil
}

func (r *ratings) PutLocalReviews(c context.Context, bookId string, rating string) error {
	_, ok := r.userAddedRatings[bookId]
	if !ok {
		r.userAddedRatings[bookId] = []string{}
	}
	r.userAddedRatings[bookId] = append(r.userAddedRatings[bookId], rating)
	r.printUserAddedRatings()
	// r.Logger().Info(r.userAddedRatings[bookId])
	return nil
}

func GetLocalReviews() {}

// CRUD ratings
func CreateReview() {}

func (r *ratings) printUserAddedRatings() {
	for k, v := range r.userAddedRatings {
		fmt.Println(k, "value is", v)
	}
}
