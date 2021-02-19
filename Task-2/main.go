package main

import (
	"encoding/json"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	fastlogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger = fastlogger.Sugar()
}

var _ validation.Validatable = Comment{}

// Comment Struct (Model)
type Comment struct {
	CommentText    string    `json:"comment"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	FavoriteNumber int       `json:"favoriteNumber"`
	GPA            float32   `json:"gpa"`
}

// ValidateComment Comment Formate
func (comment Comment) Validate() error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.CommentText, validation.Required, validation.Length(5, 50)),
		validation.Field(&comment.Email, validation.Required, is.Email),
		validation.Field(&comment.DateOfBirth, validation.Required, validation.Max(time.Now().AddDate(-12, 0, 0))),
		validation.Field(&comment.FavoriteNumber),
		validation.Field(&comment.GPA),
	)
}

func clientErrorHandler(e *json.Encoder, err error) {
	defer logger.Sync()
	logger.Error(err)
	// not ideal, but for the sake of the task it's not that bad
	errMap := map[string]string{"error": err.Error()}
	err = e.Encode(errMap)
	if err != nil {
		logger.Error(err)
	}
}

func createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var comment Comment
	jsonEncoder := json.NewEncoder(w)
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		clientErrorHandler(jsonEncoder, err)
		return
	}

	err = validation.Validate(comment)
	if err != nil {
		clientErrorHandler(jsonEncoder, err)
		return
	}

	logger.Info("Added Successfully")
	err = jsonEncoder.Encode(comment)
	if err != nil {
		logger.Error(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/comment", createComment).Methods("POST")

	addr := ":8081"
	logger.Infof("Now serving on %s", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		logger.Fatal(err)
	}
}
