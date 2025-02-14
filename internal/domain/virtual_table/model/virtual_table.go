package model

import (
	"dataset/internal/domain/datasource/model"
)

// VirtualTable is dataset item。you can read the README.md to get more information
type VirtualTable struct {
	// Id is the primary key
	Id int64
	// Name is the name of the virtual table
	// It is usually use to build the sql statement
	Name string
	// Description is the description of the virtual table
	Description string
	// CreatedAt is the created time of the virtual table
	CreatedAt int64
	// UpdatedAt is the updated time of the virtual table
	UpdatedAt int64
	// Fields is the fields of the virtual table
	//Fields []model.Field
	// DataSources is the data sources of the virtual table.
	// You can read the README.md to get more information
	DataSources []model.DataSource
}
