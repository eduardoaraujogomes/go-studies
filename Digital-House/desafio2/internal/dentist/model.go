package dentist

type Dentist struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Registration string `json:"registration"`
}

type RequestDentistCreate struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Registration string `json:"registration"`
}
type RequestDentistUpdate struct {
	FirstName    string `json:"first_name" binding:"omitempty,required"`
	LastName     string `json:"last_name" binding:"omitempty,required"`
	Registration string `json:"registration" binding:"omitempty,required"`
}
