package main

import (
	"fmt"
	"net/http"
)

func main() {
	n := &node{}
	n.addRoute("POST", "/api/user/login", Index)
	fmt.Println(n)
}

func Index(w http.ResponseWriter, req *http.Request, vars map[string]string) {
	fmt.Fprint(w, "Hello World!")
}
