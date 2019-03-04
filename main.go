package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"

	"./router"
)

func main() {
	port := "8080"
	fmt.Println("Api running on port", port)
	log.Fatal(fasthttp.ListenAndServe(":"+port, router.Instance.Handler))
}
