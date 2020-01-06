package usecase

import entity "github.com/kbsantiago/srs_assessment/pkg/customer/entity"

//Provider interface
type Provider interface {
	FindByDocument(cpf string) (*entity.Customer, error)
}

//Updater interface
type Updater interface {
	Create(customer *entity.Customer) error
}

//UseCase use case interface
type UseCase interface {
	Provider
	Updater
}
