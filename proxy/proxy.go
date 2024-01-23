package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	hostUrl, exists := os.LookupEnv("REWRITE_HOST_URL")

	if !exists {
		panic("REWRITE_URL not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received: " + r.URL.Path)

		var response *http.Response
		var err error

		if r.Method == "GET" {
			response, err = http.Get(hostUrl + r.URL.Path)
			fmt.Println("GET request sent to: " + hostUrl + r.URL.Path)
		} else {
			panic("Method not supported")
		}

		defer response.Body.Close()

		if err != nil {
			panic(err)
		}

		data, err := io.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		response.Body.Read(data)

		w.Write(data)
	})

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
