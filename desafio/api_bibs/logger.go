package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string)  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Printf("%s\t%s\t%s\t%s\tBody: %s", r.Method, r.RequestURI, name, time.Since(start), string(body))
	})
}
