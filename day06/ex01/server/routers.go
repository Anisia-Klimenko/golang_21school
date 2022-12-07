package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Try /admin endpoint")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AdminPost",
		strings.ToUpper("Post"),
		"/admin",
		Admin,
	},

	Route{
		"AdminIndex",
		"GET",
		"/admin",
		AdminIndex,
	},

	Route{
		"NewGet",
		"GET",
		"/admin/new",
		NewGet,
	},

	Route{
		"NewPost",
		"POST",
		"/admin/new",
		NewPost,
	},

	Route{
		"ShowPost",
		"GET",
		"/admin/show",
		ShowPost,
	},

	Route{
		"Logo",
		"GET",
		"/amazing_logo.png",
		SendLogo,
	},
}
