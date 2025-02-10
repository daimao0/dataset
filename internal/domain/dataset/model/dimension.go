package model

import (
	"dataset/internal/common/enum"
	"time"
)

// Dimension is a field of type dimension
type Dimension struct {

	// ID is the of the dimension
	Id int64

	// Name is the name of the field and unique within the dataset
	Name string

	// Label is the name of the field and unique within the dataset
	Label string

	// Description is the description of the  dimension
	Description string

	// DataType is the type of the dimension
	DataType enum.DataType

	// CreatedAt is the time the dimension was created
	CreatedAt time.Time

	// UpdateAt is the time the dimension was created
	UpdatedAt time.Time
}

// GetName returns the name of the field
func (m *Dimension) GetName() string {
	return m.Name
}

// GetType returns the type of the field
func (m *Dimension) GetType() enum.FieldType {
	return enum.DimensionType
}
