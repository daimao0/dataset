package model

import (
	"dataset/internal/common/enum"
)

// Field is a field of a dataset, which can be a measure, metric, dimension or driver dimension
type Field interface {

	// GetName returns the name of the field
	GetName() string

	// GetType returns the type of the field
	GetType() enum.FieldType
}
