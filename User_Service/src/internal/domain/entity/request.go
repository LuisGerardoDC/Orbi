package entity

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id,omitempty"`
}
