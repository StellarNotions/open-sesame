package router

import (
	"net/http"

	handler "github.com/StellarNotions/open-sesame/pkg/api/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetMembers",
		"GET",
		"/members",
		handler.GetMembers,
	},
	Route{
		"GetMember",
		"GET",
		"/member/{id}",
		handler.GetMember,
	},
	Route{
		"GetGates",
		"GET",
		"/gates",
		handler.GetGates,
	},
	Route{
		"GetGate",
		"GET",
		"/gate/{id}",
		handler.GetGate,
	},
}
