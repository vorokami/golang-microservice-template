package model

import "time"

type GetMethodTemplateRequest struct {
	Title    string `json:"title" binding:"required" example:"Github"`
	Category *int   `json:"category" binding:"omitempty" example:"1"`
}

type GetMethodTemplateResponse struct {
	Title     string    `json:"title"`
	Category  *int      `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
