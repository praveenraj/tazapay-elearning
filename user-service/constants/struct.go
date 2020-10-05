package constants

// NewUser data object for new user registration
type NewUser struct {
	FirstName string `valid:"Required;Match(/^.*?$/)" json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `valid:"Required;Match(/^[0-9]{10}$/)" json:"mobile"`
	Email     string `valid:"Required;Match(/^[a-zA-Z0-9_!#$%&’*+/=?{|}~^.-]+@[a-zA-Z0-9.-]+$/)" json:"email"`
}
