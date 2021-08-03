package entity

import "time"

type Measurement struct {
	ID        ID
	CreatedAt int64
	Quantity  int
}

func NewMeasurement(quantity int) (*Measurement, error) {
	m := &Measurement{
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

func (m *Measurement) Validate() error {
	if m.Quantity < 0 || m.Quantity > 1500 {
		return ErrInvalidEntity
	}

	return nil
}
