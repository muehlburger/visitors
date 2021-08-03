package measurement

import "visitors.it-zt.at/domain/entity"

// Reader interface
type Reader interface {
	//Get(id entity.ID) (*entity.Measurement, error)
	//Search(query string) ([]*entity.Measurement, error)
	//List() ([]*entity.Measurement, error)
}

// Writer write Measurements
type Writer interface {
	Create(e *entity.Measurement) (entity.ID, error)
	//Update(e *entity.Measurement) error
	//Delete(id entity.ID)
}

type Repository interface {
	Reader
	Writer
}

// UseCases define the measurement use cases
type UseCase interface {
	GetMeasurement(id entity.ID) (*entity.Measurement, error)
	SearchMeasurement(query string) ([]*entity.Measurement, error)
	ListMeasurements() ([]*entity.Measurement, error)
	CreateMeasurement(quantity int) (entity.ID, error)
	UpdateMeasurement(e *entity.Measurement) error
	DeleteMeasurement(id entity.ID) error
}
