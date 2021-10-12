package contract

// An user
//
// swagger:model ClientDTO
type ClientDTO struct {
	// The id
	//
	// required: false
	Id uint `json:"id"`
	// The name
	//
	// required: false
	Name string `json:"name"`
	// The e-mail
	//
	// required: false
	Email          string `json:"email"`
	Address        string `json:"address"`
	VATNumber      string `json:"vatNumber"`
	VATIncluded    bool   `json:"vatIncluded"`
	OrganisationId uint   `json:"organisationdid"`
}
