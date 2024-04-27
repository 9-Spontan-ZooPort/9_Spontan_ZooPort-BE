package model

type CreateSpeciesRequest struct {
	ID    string `json:"id" binding:"required,max=255"`
	Name  string `json:"name" binding:"required,max=255"`
	Class string `json:"class" binding:"required,max=255"`
}

type CreateSpeciesResponse struct {
	ID string `json:"id"`
}
