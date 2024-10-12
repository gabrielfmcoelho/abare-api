package user

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOrUpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Filiation string `json:"filiation"`
	Role      int64  `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r *CreateOrUpdateUserRequest) ValidateCreation() error {
	if r.FirstName == "" {
		return errParamIsRequired("first_name", "string")
	}
	if r.LastName == "" {
		return errParamIsRequired("last_name", "string")
	}
	if r.Gender == "" {
		return errParamIsRequired("gender", "string")
	}
	if r.Phone == "" {
		return errParamIsRequired("phone", "string")
	}
	if r.Filiation == "" {
		return errParamIsRequired("filiation", "string")
	}
	if r.Role > 0 {
		return errParamIsRequired("role", "int64")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}

func (r *CreateOrUpdateUserRequest) ValidateUpdate() error {
	if r.FirstName != "" && r.LastName != "" && r.Gender != "" && r.Phone != "" && r.Filiation != "" && r.Role != 0 && r.Email != "" && r.Password != "" {
		return nil
	}
	return fmt.Errorf("at least one field must be updated")
}
