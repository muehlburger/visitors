package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"visitors.it-zt.at/domain/entity"
)

func TestNewMeasurement(t *testing.T) {
	m, err := entity.NewReading(4)
	assert.Nil(t, err)
	assert.Equal(t, m.Quantity, 4)
	assert.NotNil(t, m.ID)
}

func TestMeasurementValidate(t *testing.T) {
	type test struct {
		quantity int
		want     error
	}

	tests := []test{
		{
			quantity: 0,
			want:     nil,
		},
		{
			quantity: 1,
			want:     nil,
		},
		{
			quantity: -1,
			want:     entity.ErrInvalidEntity,
		},
		{
			quantity: 1501,
			want:     entity.ErrInvalidEntity,
		},
		{
			quantity: 1499,
			want:     nil,
		},
		{
			quantity: 1500,
			want:     nil,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewReading(tc.quantity)
		assert.Equal(t, err, tc.want)
		assert.Equal(t, err, tc.want)
	}
}
