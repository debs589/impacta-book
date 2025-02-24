package main

import (
	"fmt"
	"impacta-book/src/router"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Running application")
	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))

}
