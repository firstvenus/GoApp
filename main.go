package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	
)
type Article struct {
	Id 		string`json:"Id"`
	Title 	string`json:"Title"`
	Desc 	string`json:"desc"`
	Content string`json:"content"`
}
var Articles []Article;

func homePage(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w, "welcome to the Homepage!")
	fmt.Println("homepage");//for writing inside ResponseWriter
}
func returnAllArticles(w http.ResponseWriter,r *http.Request){
	fmt.Println("You are at returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Id:"+key)
}
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqContent, _ := ioutil.ReadAll(r.Body)//read body,exist in body,bodyi okudum body json objesi var json olusturcam
	var art Article
	json.Unmarshal(reqContent, &art)// json objemizipars ettik içine unmarshale ile art structın içine
	Articles = append(Articles, art)//appendle ekledik
	json.NewEncoder(w).Encode(art)//encode ettink json objesine
}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/home", homePage)
	myRouter.HandleFunc("/returnAllArticles", returnAllArticles)
	myRouter.HandleFunc("/returnSingleArticles/{id}", returnSingleArticle)
	myRouter.HandleFunc("/createNewArticle", createNewArticle).Methods("POST")//yazmazsan get method olur
	myRouter.HandleFunc("/deleteArticle/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func main() {
	Articles = []Article{
		Article{Id: "1", Title: "tittle1", Desc: "desc1", Content: "content1"},
		Article{Id: "2", Title: "tittle2", Desc: "desc2", Content: "content2"},
	}
	handleRequests()
}
