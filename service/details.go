package service

import "github.com/ServiceWeaver/weaver"

type Details interface {
}

type details struct {
	weaver.Implements[Details]
}

// GET ISBN
// GET BOOK DETAILS
// GET DETAILS FROM EXTERNAL SERVICE
