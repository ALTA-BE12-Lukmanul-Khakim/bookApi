package repository

import (
	"bookapi/feature/book/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(ID uint) (domain.Basic, error) {
	var resQry Book
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Basic{}, err
	}

	if err := rq.db.Delete(&resQry).Error; err != nil {
		return domain.Basic{}, err
	}

	res := ToDomain(resQry)
	return res, nil

}

// GetAll implements domain.Repository
func (rq *repoQuery) GetAll() ([]domain.Basic, error) {
	var resQry []Book
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil

}

// GetBook implements domain.Repository
func (rq *repoQuery) GetBook(ID uint) (domain.Basic, error) {
	var resQry Book
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Basic{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

// Insert implements domain.Repository
func (rq *repoQuery) Insert(newBook domain.Basic) (domain.Basic, error) {
	var cnv Book
	cnv = FromDomain(newBook)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Basic{}, err
	}
	// selesai dari DB
	newBook = ToDomain(cnv)
	return newBook, nil
}

// Update implements domain.Repository
func (rq *repoQuery) Update(updatedBook domain.Basic) (domain.Basic, error) {
	var cnv Book
	cnv = FromDomain(updatedBook)
	if err := rq.db.Save(&cnv).Error; err != nil {
		return domain.Basic{}, err
	}
	// selesai dari DB
	updatedBook = ToDomain(cnv)
	return updatedBook, nil
}

// func (rq *repoQuery) Get(ID uint) (domain.Basic, error) {
// 	var resQry Book
// 	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
// 		return domain.Basic{}, err
// 	}
// 	// selesai dari DB
// 	res := ToDomain(resQry)
// 	return res, nil
// }
