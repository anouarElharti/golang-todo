package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anouarElharti/golang-todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server at PORT 9000...")
	log.Fatal(http.ListenAndServe(":9000", r))
}
