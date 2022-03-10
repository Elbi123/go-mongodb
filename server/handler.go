package server

import (
	"fmt"
	"net/http"
)

func RegisterAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegisterAuthor: called")
}

func FindAuthorByName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FindAuthorByName: called")
}

func FindAllAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FindAllAuthor: called")
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateAuthor: called")
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteAuthor: called")
}
