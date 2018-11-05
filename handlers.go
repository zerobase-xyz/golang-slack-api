package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Sendslack(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var slackURL = os.Getenv("slackURL") // Get Environment Variables
	sl := NewSlack(slackURL, "This is a line of text.\nAnd this is another one.", "user", ":eyes:", "", "")
	sl.Send()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func Reciveslack(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var msg Reslack
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(output)
}
