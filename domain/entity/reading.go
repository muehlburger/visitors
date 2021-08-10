package entity

import "time"

type Reading struct {
	ID        ID
	CreatedAt int64
	Quantity  int
}

func NewReading(quantity int) (*Reading, error) {
	m := &Reading{
		ID:        NewID(),
		Quantity:  quantity,
		CreatedAt: time.Now().Unix(),
	}

	err := m.Validate()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Reading) Validate() error {
	if m.Quantity < 0 || m.Quantity > 1500 {
		return ErrInvalidEntity
	}

	return nil
}
