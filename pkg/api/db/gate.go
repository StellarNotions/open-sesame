package db

import (
	model "github.com/StellarNotions/open-sesame/pkg/api/models"
)

var gates []model.Gate

// Insert allows populating database
func InsertGate(gate model.Gate) {
	gates = append(gates, gate)
}

// Get returns the whole database
func GetGate() []model.Gate {
	return gates
}
