package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	db "github.com/StellarNotions/open-sesame/pkg/api/db"
)

// IDParam is used to identify a person
//
// swagger:parameters listPerson
//type IDParam struct {
//	// The ID of a person
//	//
//	// in: path
//	// required: true
//	ID int64 `json:"id"`
//}

// GetMembers is an httpHandler for route GET /people
func GetMembers(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /members member listMembers
	//
	// Lists all members.
	//
	// This will show all recorded members.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: membersResponse
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.GetMember())
}

// GetMember is an httpHandler for route GET /people/{id}
func GetMember(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /people/{id} people listPerson
	//
	// Lists person from their id.
	//
	// This will show the record of an identified person.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//       id: IDParam
	//
	//     Responses:
	//       200: personResponse
	//       404: jsonError
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, item := range db.GetMember() {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
			return
		}
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
		panic(err)
	}
}
