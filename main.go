package main

import (
	"github.com/laidingqing/feichong/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
