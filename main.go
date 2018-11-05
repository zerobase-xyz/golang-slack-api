package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/send_slack", Logging(Sendslack, "send-slack"))
	router.POST("/recive_slack", Logging(Reciveslack, "recive-slack"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
