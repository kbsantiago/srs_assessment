package repository

import entity "github.com/kbsantiago/srs_assessment/pkg/customer/entity"

//Provider interface
type Provider interface {
	FindByDocument(cpf string) (entity.Customer, error)
}

//Updater writer
type Updater interface {
	Create(customer entity.Customer) error
}

//Repository database repository
type Repository interface {
	Provider
	Updater
}
