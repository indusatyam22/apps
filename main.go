package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	api(url)

}
func api(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(response)

}
