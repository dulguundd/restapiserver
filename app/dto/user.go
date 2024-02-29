package dto

type NewUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidationError struct {
	ValidationError NewUserValidationError `json:"validationError"`
}

type NewUserValidationError struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
