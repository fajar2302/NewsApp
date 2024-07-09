package service

import (
	articles "NEWSAPP/features/Articles"
	"errors"
)

type artikelService struct {
	artikelData articles.DataArtikelInterface
}

func New(ar articles.DataArtikelInterface) articles.ServiceArtikelInterface {
	return &artikelService{
		artikelData: ar,
	}
}

// Create implements articles.ServiceArtikelInterface.
func (ar *artikelService) Create(artikel articles.Artikel) error {
	if artikel.ArtikelName == "" {
		return errors.New("artikel name cannot be empty")
	}
	err := ar.artikelData.Insert(artikel)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements articles.ServiceArtikelInterface.
func (ar *artikelService) Delete(id uint, userid uint) error {
	if id <= 0 {
		return errors.New("invalid artikel ID")
	}
	cekuserid, err := ar.artikelData.SelectById(id)
	if err != nil {
		return err
	}

	if cekuserid.UserID != userid {
		return errors.New("user id not match, cannot delete articles")
	}

	return ar.artikelData.Delete(id)
}

// GetById implements articles.ServiceArtikelInterface.
func (ar *artikelService) GetById(id uint) (artikel *articles.Artikel, err error) {
	if id <= 0 {
		return nil, errors.New("id not valid")
	}
	return ar.artikelData.SelectById(id)
}

// GetByUserId implements articles.ServiceArtikelInterface.
func (ar *artikelService) GetAllArtikel() ([]articles.Artikel, error) {
	return ar.artikelData.GetAll()
}

// Update implements articles.ServiceArtikelInterface.
func (ar *artikelService) Update(id uint, userid uint, artikel articles.Artikel) error {
	if id == 0 {
		return errors.New("invalid artikel ID")
	}
	if artikel.ArtikelName == "" {
		return errors.New("artikel name cannot be empty")
	}

	cekuserid, err := ar.artikelData.SelectById(id)
	if err != nil {
		return err
	}

	if cekuserid.UserID != userid {
		return errors.New("user id not match, cannot update artikel")
	}

	return ar.artikelData.Update(id, artikel)
}
