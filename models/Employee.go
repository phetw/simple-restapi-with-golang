package models

// Employee struct
type Employee struct {
	ID        string `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Position  string `json:"position"`
}
