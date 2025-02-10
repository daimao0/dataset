package dto

import "time"

// DatasetDTO provide dataset info to external services
type DatasetDTO struct {

	// ID of dataset
	ID int64

	// Name of dataset
	Name string

	// Description of dataset
	Description string

	// CreatedAt of dataset
	CreatedTime time.Time

	// UpdatedAt of dataset
	UpdatedTime time.Time

	// Measures of dataset
	Measures []MeasureDTO

	// Metrics of dataset
	Metrics []MetricDTO

	// Dimensions of dataset
	Dimensions []DimensionDTO

	// DerivedDimensions of dataset
	DerivedDimensions []DriverDimensionDTO
}
