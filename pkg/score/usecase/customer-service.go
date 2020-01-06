package usecase

import (
	entity "github.com/kbsantiago/srs_assessment/pkg/customer/entity"
	repository "github.com/kbsantiago/srs_assessment/pkg/customer/repository"
)

//Service service interface
type Service struct {
	repo repository.Repository
}

//NewService create new service
func NewService(r repository.Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create an customer
func (s *Service) Create(c entity.Customer) error {
	return s.repo.Create(c)
}

//FindByDocument passing the document value
func (s *Service) FindByDocument(cpf string) (entity.Customer, error) {
	return s.repo.FindByDocument(cpf)
}
