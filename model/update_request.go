package model

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}