package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type Ratings interface {
	PutLocalReviews(context.Context, string, Review) error
	GetRatingsByBookId(context.Context, string) (rating, error)
	// HandleGetRatings(w http.ResponseWriter, r *http.Request)
}

type Review struct {
	weaver.AutoMarshal
	Reviewer string
	Stars    string
	Text     string
}

type rating struct {
	weaver.AutoMarshal
	Id string
	// Rating  []string
	Reviews []Review
}

type ratings struct {
	weaver.Implements[Ratings]
	userAddedRatings map[string][]string
	userAddedReviews map[string][]Review
}

func (r *ratings) Init(c context.Context) error {
	var err error
	r.userAddedRatings = map[string][]string{}
	r.userAddedReviews = map[string][]Review{}

	return err
}

func (r *ratings) GetRatingsByBookId(c context.Context, bookId string) (rating, error) {
	r.printUserAddedRatings()
	var err error
	bookReviews, exists := r.userAddedReviews[bookId]
	if !exists {
		err := errors.New("falha ao buscar id")
		r.Logger().Error("erro ao buscar rating com esse id!", err)
	}
	return rating{Id: bookId, Reviews: bookReviews}, err
}

func (r *ratings) PutLocalReviews(c context.Context, bookId string, rev Review) error {
	// _, ok := r.userAddedRatings[bookId]
	_, exists := r.userAddedReviews[bookId]
	if !exists {
		r.userAddedReviews[bookId] = []Review{}
	}
	r.Logger().Info("adicionando review!", rev)
	s := append(r.userAddedReviews[bookId], rev)
	r.userAddedReviews[bookId] = s

	r.Logger().Info("estado das reviews: ", r.userAddedReviews[bookId])
	return nil
}

func (r *ratings) printUserAddedRatings() {
	for k, v := range r.userAddedRatings {
		fmt.Println(k, "value is", v)
	}
}
