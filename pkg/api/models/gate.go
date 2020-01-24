package models

// Member description.
// swagger:model member
type Gate struct {
	// ID of the gate
	//
	// required: true
	ID string `json:"id,omitempty"`
	// Name of the gate
	//
	// required: true
	Name string `json:"name,omitempty"`
	// Status of gate
	//
	// required: true
	Status string `json:"status,omitempty"`
	// GPIO pin for gate
	//
	// required: true
	GPIOPin string `json:"gpioPin,omitempty"`
}
