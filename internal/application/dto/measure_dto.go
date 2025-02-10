package dto

import (
	"dataset/internal/common/enum"
	"time"
)

// MeasureDTO for dataset dto
type MeasureDTO struct {

	// Id is the id of the Measure
	Id int64

	// Name is the name of the Field and unique within the Dataset
	Name string

	// Label is the label of the Field and unique within the Dataset
	Label string

	// DataType is the data type of the Measure
	DataType enum.DataType

	// CreateTime is the time the Measure was created
	CreateTime time.Time

	// UpdateTime is the time the Measure was created
	UpdateTime time.Time
}
