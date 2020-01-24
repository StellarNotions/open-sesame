package db

import (
	model "github.com/StellarNotions/open-sesame/pkg/api/models"
)

var members []model.Member

// Insert allows populating database
func Insert(member model.Member) {
	members = append(members, member)
}

// Get returns the whole database
func Get() []model.Member {
	return members
}
