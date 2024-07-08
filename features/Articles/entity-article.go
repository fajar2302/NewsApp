package articles

type Artikel struct {
	ArtikelID      uint
	UserID         uint
	ArtikelPicture string
	ArtikelName    string
	Tag            string
	Description    string
}

type DataArtikelInterface interface {
	Insert(artikel Artikel) error
	Delete(id uint) error
	Update(id uint, artikel Artikel) error
	GetAll() ([]Artikel, error)
	SelectById(id uint) (*Artikel, error)
}

type ServiceArtikelInterface interface {
	Create(artikel Artikel) error
	Delete(id uint, userid uint) error
	Update(id uint, userid uint, artikel Artikel) error
	GetById(id uint) (artikel *Artikel, err error)
	GetAllArtikel() ([]Artikel, error)
}
