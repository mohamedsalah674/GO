package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Title 1",
			Desc:    "Description 1",
			Content: "Content 1",
		},
	}
	fmt.Println("End Point: All articles")
	json.NewEncoder(w).Encode(articles)
}

func testPostAtricles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("End Point: testPostAtricles")
	fmt.Fprintf(w, "Post testPostAtricles")
}

func homePage(w http.ResponseWriter, r *http.Request) {

	// Fprintf  --> used to print on Client side (Browser)
	fmt.Fprintf(w, "Homepage Endpoint")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostAtricles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
