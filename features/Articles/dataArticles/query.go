package dataarticles

import (
	articles "NEWSAPP/features/Articles"

	"gorm.io/gorm"
)

type artikelQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) articles.DataArtikelInterface {
	return &artikelQuery{
		db: db,
	}
}

// Delete implements articles.DataArtikelInterface.
func (a *artikelQuery) Delete(id uint) error {
	tx := a.db.Delete(&Articles{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

// Insert implements articles.DataArtikelInterface.
func (a *artikelQuery) Insert(artikel articles.Artikel) error {
	artikelGorm := Articles{
		UserID:         artikel.UserID,
		ArtikelPicture: artikel.ArtikelPicture,
		ArtikelName:    artikel.ArtikelName,
		Tag:            artikel.Tag,
		Description:    artikel.Description,
	}
	tx := a.db.Create(&artikelGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectById implements articles.DataArtikelInterface.
func (a *artikelQuery) SelectById(id uint) (*articles.Artikel, error) {
	var artikelGorm Articles
	tx := a.db.First(&artikelGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping
	var projectcore = articles.Artikel{
		ArtikelID:   artikelGorm.ID,
		UserID:      artikelGorm.UserID,
		ArtikelName: artikelGorm.ArtikelName,
		Tag:         artikelGorm.Tag,
		Description: artikelGorm.Description,
	}

	return &projectcore, nil
}

// SelectByUserId implements articles.DataArtikelInterface.
func (a *artikelQuery) GetAll() ([]articles.Artikel, error) {
	var allArtikel []Articles // var penampung data yg dibaca dari db
	tx := a.db.Find(&allArtikel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allArtikelCore []articles.Artikel
	for _, v := range allArtikel {
		allArtikelCore = append(allArtikelCore, articles.Artikel{
			UserID:         v.UserID,
			ArtikelPicture: v.ArtikelPicture,
			ArtikelName:    v.ArtikelName,
			Tag:            v.Tag,
			Description:    v.Description,
		})
	}

	return allArtikelCore, nil
}

// Update implements articles.DataArtikelInterface.
func (a *artikelQuery) Update(id uint, artikel articles.Artikel) error {
	var artikelGorm Articles
	tx := a.db.First(&artikelGorm, id)
	if tx.Error != nil {
		return tx.Error
	}
	artikelGorm.ArtikelPicture = artikel.ArtikelPicture
	artikelGorm.ArtikelName = artikel.ArtikelName
	artikelGorm.Tag = artikel.Tag
	artikelGorm.Description = artikel.Description

	tx = a.db.Save(&artikelGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
