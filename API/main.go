package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for Article -file

type Article struct {
	ID      int     `json:"article_id"`
	Title   string  `json:"article_title"`
	Content string  `json:"article_content"`
	Author  *Author `json:"author"`
}

type Author struct {
	ID       string `json:"author_id"`
	Fullname string `json:"fullname"`
}

// demo DB
var Articles []Article

// middleware , helper - file
func (c *Article) validate() bool {
	if c.Title == "" || c.Content == "" || c.Author == nil {
		return false
	}
	return true
}

// controllers - file

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Articles API</h1>"))
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	// articlesJson, _ := json.Marshal(Articles)
	w.Header().Set("Content-Type", "application/json")
	// w.Write(articlesJson)
	json.NewEncoder(w).Encode(Articles)
}

func getArticleByID(w http.ResponseWriter, r *http.Request) {
	// get the article id from the request path
	// id := r.URL.Query().Get("id")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	// loop through articles and find the matching id

	for _, article := range Articles {
		if article.ID == int(id) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	http.Error(w, "Article not found", http.StatusNotFound)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		http.Error(w, "There is no data", http.StatusBadRequest)
		return
	}
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	if !article.validate() {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}
	randomID := rand.Intn(1000)
	article.ID = randomID
	Articles = append(Articles, article)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	prams := mux.Vars(r)
	id, _ := strconv.ParseInt(prams["id"], 10, 64)
	for index, article := range Articles {
		if article.ID == int(id) {
			Articles = append(Articles[:index], Articles[index+1:]...)
			var updatedArticle Article
			_ = json.NewDecoder(r.Body).Decode(&updatedArticle)
		}
	}
	w.WriteHeader(http.StatusOK)
}

func deleteArticale(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for i, artical := range Articles {
		if artical.ID == int(id) {
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}
	w.WriteHeader(http.StatusNoContent)

}

func main() {
	r := mux.NewRouter()
	// seeding some data
	Articles = append(Articles, Article{ID: 1, Title: "First Article", Content: "This is the content of the first article", Author: &Author{ID: "1", Fullname: "Wael"}})
	Articles = append(Articles, Article{ID: 2, Title: "Second Article", Content: "This is the content of the second article", Author: &Author{ID: "2", Fullname: "Ali"}})
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/articles", getAllArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", getArticleByID).Methods("GET")
	r.HandleFunc("/articles", createArticle).Methods("POST")
	r.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", deleteArticale).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
