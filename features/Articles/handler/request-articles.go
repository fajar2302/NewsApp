package handler

type ArtikelRequest struct {
	ArtikelName string `json:"articles_name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}
