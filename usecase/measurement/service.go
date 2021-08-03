package measurement

import (
	"visitors.it-zt.at/domain/entity"
)

type Service struct {
	repo Repository
}

// NewService creates a new measurement service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateMeasurement creates a measurement
func (s *Service) CreateMeasurement(quantity int) (entity.ID, error) {
	m, err := entity.NewMeasurement(quantity)
	if err != nil {
		return m.ID, err
	}
	return s.repo.Create(m)
}
