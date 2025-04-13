package Request

type RegisterRequest struct {
	Name      string `json:"name" binding:"required"`
	Document  string `json:"document" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
}

type ClinicalRegisterRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}

type AddressRequest struct {
	ZipCode    string `json:"zipCode"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	City       string `json:"city"`
	State      string `json:"state"`
}
