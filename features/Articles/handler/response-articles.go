package handler

type ArtikelResponse struct {
	UserID         uint   `json:"user_id"`
	ArtikelPicture string `json:"artikel_picture"`
	ArtikelName    string `json:"articles_name"`
	Tag            string `json:"tag"`
	Description    string `json:"description"`
}
