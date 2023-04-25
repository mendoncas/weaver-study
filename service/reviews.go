package service

import "github.com/ServiceWeaver/weaver"

type Reviews interface {
}

type reviews struct {
	weaver.Implements[Reviews]
}

// getRatings
// bookReviewsByID
