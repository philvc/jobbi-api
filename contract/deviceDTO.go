package contract

// An user
//
// swagger:model DeviceDTO
type DeviceDTO struct {
	// The id
	//
	// required: true
	Id string `json:"id"`
	// The token
	//
	// required: false
	Token string `json:"token"`
}
