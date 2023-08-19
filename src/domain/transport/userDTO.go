package transport

type UserDTO struct {
	Email    string `json:"email"`
	PassHash string `json:"passHash"`
}
