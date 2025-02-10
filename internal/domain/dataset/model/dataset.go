package model

import (
	"dataset/internal/domain/data_source/model"
	"time"
)

// Dataset provides field metadata information such as measure, metric, dimension, and derived dimension, and also supplies data.
type Dataset struct {
	// ID of dataset
	Id int64

	// Name of dataset
	Name string

	// Description of dataset
	Description string

	// CreatedAt of dataset
	CreatedAt time.Time

	// UpdatedAt of dataset
	UpdatedAt time.Time

	// Fields of dataset contains measure、metric、dimension、derived dimension
	Fields []Field

	// DataSource uses the data source to read data
	DataSource model.DataSource
}
