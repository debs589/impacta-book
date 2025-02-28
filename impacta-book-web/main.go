package main

import (
	"fmt"
	"impacta-book/src/router"
	"impacta-book/src/utils"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Running application")
	utils.LoadTemplates()

	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))

}
