package main

import (
	"fmt"
	"net/http"
)

func main() {
	n := &node{}
	n.insertChild("POST", "/api/user/register", Index)
	n.insertChild("GET", "/api/user/register", Index)
	fmt.Println(n)
}

func Index(w http.ResponseWriter, req *http.Request, vars map[string]string) {
	fmt.Fprint(w, "Hello World!")
}
