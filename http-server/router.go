package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	apiV1Router := router.
		PathPrefix("/api/v1").
		Subrouter()
	apiV1Router.HandleFunc("/hello-world", helloWorld)

	apiV2Router := router.
		PathPrefix("/api/v2").
		Subrouter()
	apiV2Router.
		HandleFunc("/headers", headers).
		Methods("GET", "POST")
	apiV2Router.
		Path("/params/{type:[a-zA-Z]+}/{id:[0-9]+}").
		Methods("GET", "POST").
		HandlerFunc(pathParameters)
	apiV2Router.
		Path("/query").
		Methods("GET").
		Queries("type", "{type}", "id", "{id:[0-9]+}").
		HandlerFunc(queryStrings)
	apiV2Router.
		Path("/json-body").
		Methods("POST").
		HandlerFunc(jsonBody)

	return router
}
