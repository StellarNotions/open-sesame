package handlers

import (
	"encoding/json"
	"github.com/StellarNotions/open-sesame/pkg/api/db"
	"github.com/gorilla/mux"
	"net/http"
)

// GetGates is an httpHandler for route GET /gates
func GetGates(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /gates gate getGates
	//
	// Gets all gates
	//
	// Lists all gates
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
	//       200: gatesResponse
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.GetGate())
}

// GetGate is an httpHandler for route GET /gate/{id}
func GetGate(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /gate/{id} gate listGate
	//
	// Lists gate from it's id.
	//
	// This will show the record of an identified gate.
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
	//       200: gateResponse
	//       404: jsonError
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, item := range db.GetGate() {
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

//// OpenGate is an httpHandler for route POST /gate/open
//func OpenGate(w http.ResponseWriter, r *http.Request) {
//	// swagger:route POST /gate/open gate openGate
//	//
//	// Open gate
//	//
//	// This will cause the gate to open
//	//
//	//     Consumes:
//	//     - application/json
//	//
//	//     Produces:
//	//     - application/json
//	//
//	//     Schemes: http, https
//	//
//	//     Responses:
//	//       200: gateOpenedResponse
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(db.GetGate())
//}
