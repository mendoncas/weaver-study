package service

import "github.com/ServiceWeaver/weaver"

type Ratings interface {
}

type ratings struct {
	weaver.Implements[Ratings]
}

//CRUD ratings
