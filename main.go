package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mendoncas/weaver-study/service"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	// Initialize the Service Weaver application.
	root := weaver.Init(context.Background())
	r := mux.NewRouter()

	// books, err := weaver.Get[service.Books](root)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// reverser, err := weaver.Get[service.Reverser](root)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	type ratingDto struct {
		id     int
		rating int
	}

	details, err := weaver.Get[service.Details](root)
	if err != nil {
		log.Fatal(err)
	}

	reviews, err := weaver.Get[service.Reviews](root)
	if err != nil {
		log.Fatal(err)
	}

	ratings, err := weaver.Get[service.Ratings](root)
	if err != nil {
		log.Fatal(err)
	}

	// Get a network listener on address "localhost:12345".
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}

	r.HandleFunc("/details/{id}", func(w http.ResponseWriter, r *http.Request) {
		// details.AddBookDetails(r.Context(), "aaa")
		res, _ := details.FindDetailsByBookId(r.Context(), mux.Vars(r)["id"])
		rp, _ := json.Marshal(res)
		w.Write(rp)
	}).Methods("GET")

	r.HandleFunc("/reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		reviews.GetAllBookReviews(r.Context(), "aa")
	})

	r.HandleFunc("/ratings/{id}", func(w http.ResponseWriter, r *http.Request) {
		rtg, _ := ratings.GetRatingsByBookId(r.Context(), mux.Vars(r)["id"])
		fmt.Printf("rtg: %v\n", rtg)
		json.NewEncoder(w).Encode(rtg)

	}).Methods("GET")

	r.HandleFunc("/ratings", func(w http.ResponseWriter, r *http.Request) {
		var bookRating ratingDto
		err := json.NewDecoder(r.Body).Decode(&bookRating)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("dados do json:")
		fmt.Println(bookRating.id, bookRating.rating)
		fmt.Fprintf(w, "%+v", bookRating)
		ratings.PutLocalReviews(r.Context(), strconv.Itoa(bookRating.id), strconv.Itoa(5))
	}).Methods("POST")

	http.Serve(lis, r)
}
