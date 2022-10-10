package repository

import (
	"bookapi/feature/user/domain"

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

// Add User
func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newUser = ToDomain(cnv)
	return newUser, nil
}

// GetUser implements domain.Repository
func (rq *repoQuery) GetUser(Nama string, Password string) (domain.Core, error) {
	var resQry User
	if err := rq.db.Where("Nama = ? and Password =?", Nama, Password).First(&resQry).Error; err != nil {
		return domain.Core{}, err
	}

	loginUser := ToDomain(resQry)
	return loginUser, nil
}

// Update Data User
func (rq *repoQuery) Update(updatedData domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(updatedData)
	if err := rq.db.Save(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	updatedData = ToDomain(cnv)
	return updatedData, nil
}

// Ambil ID user
func (rq *repoQuery) Get(ID uint) (domain.Core, error) {
	var resQry User
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

// Ambil semua Data
func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var resQry []User
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(deleteUser domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(deleteUser)
	if err := rq.db.Delete(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	deleteData := ToDomain(cnv)
	return deleteData, nil
}
