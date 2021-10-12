package contract

// An user
//
// swagger:model DeviceDTO
type DeviceDTO struct {
	// The id
	//
	// required: true
	Id uint `json:"id"`
	// The token
	//
	// required: false
	Token string `json:"token"`

}
