package main

import (
	"fmt"
	"os"
	"log"
	"net/http"

	"github.com/MariaGiurgiu/react-project/router"
)

func main() {
	fmt.Println("Starting http server...")

	httpAddr := os.Getenv("HTTP_ADDR")
	log.Fatalf("server stopped because: %s\n", http.ListenAndServe(httpAddr, router.New()))
}