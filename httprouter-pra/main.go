package main

import (
	"fmt"
	"net/http"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func main() {
	router := New()

	router.POST("/api/login", Index)

	w := new(mockResponseWriter)

	req, _ := http.NewRequest("POST", "/api/login", nil)
	router.ServeHTTP(w, req)
}

func Index(w http.ResponseWriter, req *http.Request, vars map[string]string) {
	fmt.Println("Get Index")
	fmt.Fprint(w, "Hello World!")
}

func Register(w http.ResponseWriter, req *http.Request, vars map[string]string) {
	fmt.Fprint(w, "Hello World!")
}
