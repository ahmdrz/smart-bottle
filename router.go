package main

import (
	"net/http"

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
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"newRegisteration",
		"POST",
		"/register/new/",
		NewRegisteration,
	},
	Route{
		"validateRegistration",
		"GET",
		"/register/validate/",
		ValidateRegistration,
	},
	Route{
		"setProfile",
		"POST",
		"/profiles/set/",
		SetProfile,
	},
	Route{
		"getProfile",
		"GET",
		"/profiles/get/",
		GetProfile,
	},
	Route{
		"getNotification",
		"GET",
		"/notification/",
		GetLastNotification,
	},
	Route{
		"newDrinkRecord",
		"POST",
		"/drink/set/",
		SetDrink,
	},
}
