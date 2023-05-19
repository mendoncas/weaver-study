package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	// reviews, err := weaver.Get[service.Reviews](root)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rtg, err := weaver.Get[service.Ratings](root)
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

	// r.HandleFunc("/reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	reviews.GetAllBookReviews(r.Context(), "aa")
	// })

	r.HandleFunc("/reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		res, _ := rtg.GetRatingsByBookId(r.Context(), mux.Vars(r)["id"])
		// rp, _ := json.Marshal(res)
		// w.Write(rp)
		json.NewEncoder(w).Encode(res)
	}).Methods("GET")

	r.HandleFunc("/reviews/{id}", func(w http.ResponseWriter, r *http.Request) {
		var bookReview service.Review
		err := json.NewDecoder(r.Body).Decode(&bookReview)
		if err != nil {
			fmt.Println(err.Error())
		}
		rtg.PutLocalReviews(r.Context(), mux.Vars(r)["id"], bookReview)
		fmt.Printf("bookReview: %v\n", bookReview)
		json.NewEncoder(w).Encode(bookReview)
	}).Methods("POST")

	http.Serve(lis, r)
}
