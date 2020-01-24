//go:generate swagger generate spec

// Package classification People API.
//
// the purpose of this application is to provide an application that allows remote use of a wireless garage door opener
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:8000
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Wyatt Barnes <wyatt@writerof.software> http://opensesame.dev
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"

	"github.com/StellarNotions/open-sesame/pkg/api/db"
	model "github.com/StellarNotions/open-sesame/pkg/api/models"
	"github.com/StellarNotions/open-sesame/pkg/api/router"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {
	// populate our test database
	db.Insert(model.Member{ID: "1", FirstName: "James", LastName: "Holden", Created: time.Now().UTC().String(), Status: "active", Pin: "1234"})
	db.Insert(model.Member{ID: "2", FirstName: "Naomi", LastName: "Nagata", Created: time.Now().UTC().String(), Status: "active", Pin: "4321"})

	appRouter := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", setupGlobalMiddleware(appRouter)))
}
