package models

// Member description.
// swagger:model member
type Member struct {
	// ID of the member
	//
	// required: true
	ID string `json:"id,omitempty"`
	// First name of the member
	//
	// required: true
	FirstName string `json:"firstName,omitempty"`
	// Last Name of the member
	//
	// required: true
	LastName string `json:"lasName,omitempty"`
	// UTC time stamp of when member was created
	//
	// required: true
	Created string `json:"created,omitempty"`
	// Status of member
	//
	// required: true
	Status string `json:"status,omitempty"`
	// PIN for member
	//
	// required: true
	Pin string `json:"pin,omitempty"`
}
