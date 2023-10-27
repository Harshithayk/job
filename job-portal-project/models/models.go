package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"-" validate:"required"`
}

type Job struct {
	gorm.Model
	JobTitle  string  `json:"job_title" validate:"required"`
	JobSalary string  `json:"job_salary" validate:"required"`
	Company   Company `gorm:"ForeignKey:uid"`
	Uid       uint    `JSON:"uid, omitempty"`
}

type Company struct {
	gorm.Model
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type UserSignUp struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Userlogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type CreateJob struct {
	JobTitle  string `json:"job_title" validate:"required"`
	JobSalary string `json:"job_salary" validate:"required"`
}
