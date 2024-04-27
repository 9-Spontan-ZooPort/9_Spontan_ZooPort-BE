package model

type CreateAnimalRequest struct {
	SpeciesID       string  `json:"species_id" binding:"required,max=255"`
	Nickname        string  `json:"nickname" binding:"required,max=255"`
	BirthDate       string  `json:"birth_date" binding:"required"`
	Gender          string  `json:"gender" binding:"required,oneof=jantan betina"`
	Weight          float64 `json:"weight" binding:"required,min=0"`
	PhotoUrl        string  `json:"photo_url" binding:"required,max=255,url"`
	StatusKesehatan string  `json:"status_kesehatan" binding:"required,oneof=sehat pemulihan sakit kritis"`
	Description     string  `json:"description" binding:"required,max=65535"`
}

type CreateAnimalResponse struct {
	ID string `json:"id"`
}
