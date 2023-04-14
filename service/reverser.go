package service

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	"github.com/mendoncas/weaver-study/client"
)

// Reverser component.
type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

// Implementation of the Reverser component.
type reverser struct {
	weaver.Implements[Reverser]
	mongo client.DbClient
}

func (r *reverser) Init(context.Context) error {
	var err error
	r.mongo, err = weaver.Get[client.DbClient](r)
	return err
}

func (r *reverser) Reverse(c context.Context, s string) (string, error) {
	runes := []rune(s)
	n := len(runes)
	log := r.Logger()
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	res, err := r.mongo.Insert(c, "Runas", "Invertidas", string(runes))
	if err != nil {
		log.Error("falha ao inserir string no banco!", err)
	}
	log.Info(res)
	return string(runes), nil
}
