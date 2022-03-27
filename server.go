package main

import (
	"context"
	"fmt"
	"getMethods"
	"log"
	"net/http"
	"postMethods"
	"regexp"
	"strings"
)

type ctxKey struct{}

func main() {
	http.HandleFunc("/", Serve)
	//http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}

func withoutPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Path not called")
}

var routes = []route{
	newRoute("GET", "/", withoutPath),
	newRoute("GET", "/getAdvice", getMethods.GetAdvice),
	newRoute("POST", "/getAdviceById", postMethods.GetAdviceById),
	newRoute("GET", "/getAdviceFromDb", getMethods.GetAdviceFromDb),
	newRoute("POST", "/deleteAdvice", postMethods.DeleteAdviceDb),
	newRoute("POST", "/insertAdvice", postMethods.InsertAdviceDb),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
