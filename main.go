package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homepage")
	fmt.Println("Endppoint Hit: homepage")
}

// func handleRequest() {
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/articles", returnAllArticles)
// 	log.Fatal(http.ListenAndServe(":10000", nil))
// }

func handleRequestwithmux() {
	// creates a new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", returnAllArticles)
	// NOTE: Ordering is important here! This has to be defined before
	// the other `/article` endpoint.
	myRouter.HandleFunc("/articles", createNewArticles).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")

	myRouter.HandleFunc("/articles/{id}", returnOneArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnOneArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit singleArticles")
	vars := mux.Vars(r)
	key := vars["id"]
	// fmt.Fprintf(w, "key: "+key)

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		} else {
			fmt.Fprintf(w, "Not found")
		}
	}
}

func createNewArticles(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// get the params using mux
	vars := mux.Vars(r)
	// get the id

	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			fmt.Fprintf(w, "articl deleted")

		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(reqBody, id)

	for index, article := range Articles {
		if article.Id == id {

		}
	}
}
func main() {
	fmt.Println("Rest API V2.0 - muc Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article description here", Content: "THis is the content of this article"},
		Article{Id: "2", Title: "jamb it af call of asuu", Desc: "jamb say dem don end ASSU", Content: "The matter long sotay, he reach kanfasha, jamb and assu dey test who be oga"},
	}
	// handleRequest()
	handleRequestwithmux()
}

// Article structure is ...
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Article structure is ...
var Articles []Article
