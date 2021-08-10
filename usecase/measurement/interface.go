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
	Create(e *entity.Reading) (entity.ID, error)
	//Update(e *entity.Measurement) error
	//Delete(id entity.ID)
}

type Repository interface {
	Reader
	Writer
}

// UseCases define the measurement use cases
type UseCase interface {
	GetReading(id entity.ID) (*entity.Reading, error)
	SearchReading(query string) ([]*entity.Reading, error)
	ListReadings() ([]*entity.Reading, error)
	CreateReading(quantity int) (entity.ID, error)
	UpdateReading(e *entity.Reading) error
	DeleteReading(id entity.ID) error
}
