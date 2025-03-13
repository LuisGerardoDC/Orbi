package entity

type Response struct {
	Succes  bool   `json:"succes,omitempty"`
	Message string `json:"message,omitempty"`
	User    User   `json:"user,omitempty"`
}
