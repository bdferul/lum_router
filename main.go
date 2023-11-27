package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HandleFunc = func(w http.ResponseWriter, r *http.Request)

type Lum struct {
	Get  HandleFunc
	Post HandleFunc
}

func (l Lum) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upperMethod := strings.ToUpper(r.Method)
	methodMap := map[string]HandleFunc{
		"GET":  l.Get,
		"POST": l.Post,
	}

	f, _ := methodMap[upperMethod]
	if f != nil {
		f(w, r)
	}
}

func main() {
	addr := "0.0.0.0:1971"
	lumMap := map[string]Lum{
		"/": {
			Get: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, r.URL.Path)
			},
		},
	}

	for k, v := range lumMap {
		http.Handle(k, v)

		nameMap := map[string]HandleFunc{
			"Get":  v.Get,
			"Post": v.Post,
		}
		availableMethods := []string{}
		for methodKey, methodFunc := range nameMap {
			if methodFunc != nil {
				availableMethods = append(availableMethods, methodKey)
			}
		}
		methodsJoined := strings.Join(availableMethods, " ")
		fmt.Printf("%s: %s\n", k, methodsJoined)
	}
	fmt.Printf("Now serving at http://%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
