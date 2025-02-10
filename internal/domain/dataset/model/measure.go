package model

import (
	"dataset/internal/common/enum"
	"time"
)

// Measure is a field of type measure
type Measure struct {
	// Id is the id of the Measure
	Id int64

	// Name is the name of the Field and unique within the Dataset
	Name string

	// Label is the label of the Field and unique within the Dataset
	Label string

	// DataType is the data type of the Measure
	DataType enum.DataType

	// CreatedAt is the time the Measure was created
	CreatedAt time.Time

	// UpdatedAt is the time the Measure was created
	UpdatedAt time.Time
}

// GetName returns the name of the Measure
func (m *Measure) GetName() string {
	return m.Name
}

// GetType returns the type of the Measure
func (m *Measure) GetType() enum.FieldType {
	return enum.MeasureType
}
