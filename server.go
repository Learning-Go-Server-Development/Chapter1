package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "3000"
	msg := "Server starting on port "
	fmt.Println(msg + port)
	http.ListenAndServe(":3000", nil)
}
