package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("[로그] 방금 요청을 받았다.")
	target := os.Getenv("TARGET")
	version := "1.0.4"
	if target == "" {
		target = "세상아."
	}
	fmt.Fprintf(w, "Hello %s! ver[%s]\n", target, version)
}

func main() {
	log.Print("[로그] 서버를 시작하겠다.")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
