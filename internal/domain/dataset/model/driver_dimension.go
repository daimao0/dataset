package model

import (
	"dataset/internal/common/enum"
	"time"
)

// DriverDimension is a field of type driverDimension
type DriverDimension struct {
	// ID is the unique identifier of the driver dimension
	Id int64
	// Name is the name of the field
	Name string
	// Name is the name of the field
	Label string
	// DataType is the data type of the driverDimension
	DataType enum.DataType
	// CreatedAt is the time when the driver dimension was created
	CreatedAt time.Time
	// UpdatedAt is the time when the driver dimension was updated
	UpdatedAt time.Time
}

// GetName returns the name of the field
func (d *DriverDimension) GetName() string {
	return d.Name
}

// GetType returns the type of the field
func (d *DriverDimension) GetType() enum.FieldType {
	return enum.DriverDimensionType
}
