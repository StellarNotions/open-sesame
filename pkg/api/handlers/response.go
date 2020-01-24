package handlers

import model "github.com/StellarNotions/open-sesame/pkg/api/models"

// JsonError is a generic error in JSON format
//
// swagger:response jsonError
type jsonError struct {
	// in: body
	Message string `json:"message"`
}

// MemberResponse contains a single member's information
//
// swagger:response memberResponse
type memberResponse struct {
	// in: body
	Payload *model.Member `json:"member"`
}

// MembersResponse contains all member's information
//
// swagger:response membersResponse
type membersResponse struct {
	// in: body
	Payload *[]model.Member `json:"members"`
}
