package model

import (
	"dataset/internal/common/enum"
	"time"
)

// Metric is a field of type metric
type Metric struct {
	Id int64

	Name string

	Label string

	DataType enum.DataType

	CreatedAt time.Time

	UpdatedAt time.Time
}

// GetName returns the name of the field
func (m *Metric) GetName() string {
	return m.Name
}

// GetType returns the type of the field
func (m *Metric) GetType() enum.FieldType {
	return enum.MetricType
}
