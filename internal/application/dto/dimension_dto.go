package dto

import (
	"dataset/internal/common/enum"
	"time"
)

// DimensionDTO for dataset dto
type DimensionDTO struct {

	// Id is the id of the DimensionDTO
	Id int64

	// Name is the name of the Field and unique within the Dataset
	Name string

	// Label is the label of the Field and unique within the Dataset
	Label string

	// DataType is the data type of the DimensionDTO
	DataType enum.DataType

	// CreateTime is the time the Metric was created
	CreateTime time.Time

	// UpdateTime is the time the Metric was created
	UpdateTime time.Time
}
