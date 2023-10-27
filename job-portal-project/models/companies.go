package models

import (
	"context"
)

func (s *Conn) CreateCompany(ctx context.Context, ni CreateCompany, userId uint) (Company, error) {
	com := Company{
		CompanyName: ni.CompanyName,
		Adress:      ni.Adress,
		Domain:      ni.Domain,
	}

	tx := s.db.WithContext(ctx).Create(&com)

	if tx.Error != nil {
		return Company{}, tx.Error
	}

	return com, nil
}

func (s *Conn) ViewCompany(ctx context.Context, userId string) ([]Company, error) {
	var com []Company
	tx := s.db.Find(&com)
	err := tx.Find(&com).Error
	if err != nil {
		return nil, err
	}

	return com, err

}
func (s *Conn) Getcompany(id int64) (Company, error) {
	var m Company
	tx := s.db.Where("id=?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return Company{}, err
	}
	return m, nil

}

func (s *Conn) JobCreation(ctx context.Context, ni CreateJob, id int64) (CreateJob, error) {

	err := s.db.Create(&ni).Error
	if err != nil {
		return CreateJob{}, err
	}
	return ni, nil
}
