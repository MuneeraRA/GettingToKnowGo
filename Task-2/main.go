package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	validation "github.com/ozzo-validation"
	"github.com/ozzo-validation/is"
)

// Comment Struct (Model)
type Comment struct {
	CommentText    string    `json:"comment_text"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	FavoriteNumber int       `json:"favorite_number"`
	GPA            float32   `jaon:"gpa"`
}

// ValidateComment Comment Formate
func (comment Comment) ValidateComment() error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.CommentText, validation.Required, validation.Length(5, 50)),
		validation.Field(&comment.Email, validation.Required, is.Email),
		validation.Field(&comment.DateOfBirth, validation.Required, validation.Max(time.Now().AddDate(-12, 0, 0))),
		validation.Field(&comment.FavoriteNumber, is.Digit),
		validation.Field(&comment.GPA, is.Float),
	)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var comment Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)
	err := comment.ValidateComment()
	if err == nil {
		fmt.Println("Added Successfully")
		json.NewEncoder(w).Encode(comment)
	} else {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
	}

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/comment", createComment).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", r))
}
