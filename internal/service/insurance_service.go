package service

import (
	"github.com/cabralfbenja/segurointeligente/internal/dtos"
	"github.com/cabralfbenja/segurointeligente/internal/entities"
	"github.com/cabralfbenja/segurointeligente/internal/repository"
)

type InsuranceService interface {
	Create(dtos.InsuranceDto) (int64, error)
	GetAll(userId int64) ([]*entities.Insurance, error)
	Update(id int64) (*entities.Insurance, error)
}

type insuranceService struct {
	repo repository.InsuranceRepository
}

func NewRuleService(repo repository.InsuranceRepository) InsuranceService {
	return &insuranceService{
		repo: repo,
	}
}

func (s *insuranceService) Create(input dtos.InsuranceDto) (int64, error) {
	item, err := entities.NewInsurance(1, input)

	if err != nil {
		return 0, err
	}

	err = s.repo.Insert(item)
	return item.ID, err
}

func (s *insuranceService) GetAll(userId int64) ([]*entities.Insurance, error) {
	items, err := s.repo.GetAllByUserID(userId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *insuranceService) Update(id int64) (*entities.Insurance, error) {
	item, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
