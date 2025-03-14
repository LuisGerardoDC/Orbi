package entity

type Message struct {
	User   User   `json:"user"`
	Action string `json:"action"`
}
