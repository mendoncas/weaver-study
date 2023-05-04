package main

import (
	"context"
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

	books, err := weaver.Get[service.Books](root)
	if err != nil {
		log.Fatal(err)
	}

	reverser, err := weaver.Get[service.Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	details, err := weaver.Get[service.Details](root)
	if err != nil {
		log.Fatal(err)
	}

	reviews, err := weaver.Get[service.Reviews](root)
	if err != nil {
		log.Fatal(err)
	}

	// Get a network listener on address "localhost:12345".
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("hello listener available on %v\n", lis)

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		s, _ := reverser.Reverse(r.Context(), r.URL.Query().Get("name"))
		fmt.Fprintf(w, s)
	}).Methods("GET")

	r.HandleFunc("/details", func(w http.ResponseWriter, r *http.Request) {
		details.AddBookDetails(r.Context(), "aaa")
	})

	r.HandleFunc("/reviews", func(w http.ResponseWriter, r *http.Request) {
		reviews.GetAllBookReviews(r.Context(), "aa")
	})

	r.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tentando aqui chamar a função")
		err := books.RegisterBook(r.Context())
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	}).Methods("POST")

	http.Serve(lis, r)
}

// func HandleHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, $s\n", r.URL.Query().Get("name"))
// }
