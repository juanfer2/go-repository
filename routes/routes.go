package routes

import (
	"fmt"
	"net/http"
)

type Router struct{
	rules: map[string]http.HandlerFunc,
}

func CreateRouter() *Router  {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(w, "HI WORLD")
}