package model

type TodosResponse struct {
	Message string         `json:"message,omitempty"`
	Todos   map[int]string `json:"message,omitempty"`
	Error   string         `json:"message,omitempty"`
}
