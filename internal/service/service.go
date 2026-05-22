package service

import (
	"fmt"

	"github.com/ArtemHvozdov/task-cli/internal/entity"
)

type Repository interface {
	Add(description string) (uint16, error)
	Update(id uint16, desscription string) error
	Delete(id uint16) error
	UpdateStatus(id uint16, newStatus entity.Status) error
	GetAll() ([]entity.Task, error)
	GetByStatus(status entity.Status) ([]entity.Task, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Add(description string) (uint16, error) {
	if description == "" {
		return 0, fmt.Errorf("description cannot be empty")
	}
	return s.repo.Add(description)
}

func (s *Service) Update(id uint16, desscription string) error {
	if desscription == "" || id == 0 {
		return fmt.Errorf("description cannot be empty and id must be greater than 0")
	}
	return s.repo.Update(id, desscription)
}

func (s *Service) Delete(id uint16) error {
	if id == 0 {
		return fmt.Errorf("id must be greater than 0")
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateStatus(id uint16, newStatus entity.Status) error {
	if id == 0 {
		return fmt.Errorf("id must be greater than 0")
	}
	
	if newStatus != entity.Todo && newStatus != entity.InProgress && newStatus != entity.Done {
		return fmt.Errorf("invalid status: %s", newStatus)
	}
	
	return s.repo.UpdateStatus(id, newStatus)
}

func (s *Service) GetAll() ([]entity.Task, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByStatus(status entity.Status) ([]entity.Task, error) {
	return s.repo.GetByStatus(status)
}