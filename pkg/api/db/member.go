package db

import (
	model "github.com/StellarNotions/open-sesame/pkg/api/models"
)

var members []model.Member

// InsertMember allows populating database
func InsertMember(member model.Member) {
	members = append(members, member)
}

// GetMember returns the whole database
func GetMember() []model.Member {
	return members
}
