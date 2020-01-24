package handlers

import (
	"encoding/json"
	"github.com/StellarNotions/open-sesame/pkg/api/db"
	"github.com/StellarNotions/open-sesame/pkg/api/models"
	"github.com/StellarNotions/open-sesame/pkg/gpio/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func findGate(gates []models.Gate, id string) (bool, models.Gate) {
	for _, gate := range gates {
		if gate.ID == id {
			return true, gate
		}

		return false, models.Gate{}
	}
	return false, models.Gate{}
}

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
	ok, gate := findGate(db.GetGate(), params["id"])

	if ok {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(gate); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
		panic(err)
	}
}

// OpenGate is an httpHandler for route POST /gate/{id}/open
func OpenGate(w http.ResponseWriter, r *http.Request) {
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
	ok, gate := findGate(db.GetGate(), params["id"])

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonError{Message: "not found"}); err != nil {
			panic(err)
		}
	}

	gateOpened := handlers.OpenCloseGate(gate.GPIOPin)

	if !gateOpened {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonError{Message: "unable to open gate"}); err != nil {
			panic(err)
		}
	}

	gate.Status = "opened"

	w.WriteHeader(http.StatusOK)
	return
}

// CloseGate is an httpHandler for route POST /gate/{id}/close
func CloseGate(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /gate/{id}/close gate closeGate
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
	ok, gate := findGate(db.GetGate(), params["id"])

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonError{Message: "not found"}); err != nil {
			panic(err)
		}
	}

	gateOpened := handlers.OpenCloseGate(gate.GPIOPin)

	if !gateOpened {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(jsonError{Message: "unable to close gate"}); err != nil {
			panic(err)
		}
	}

	gate.Status = "closed"

	w.WriteHeader(http.StatusOK)
	return
}
