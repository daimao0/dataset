package model

import (
	"dataset/internal/common/enum"
	"time"
)

// Field is a field of a virtual table
type Field struct {
	// Id is the id of the field
	Id int64
	// Name is the name of the field, it must be unique in a physical table
	// if the physical table is created by sql, the name is the column name in the sql
	// if the physical table is created by data, the name is the column name in the data
	Name string
	// Label is the label of the field,it must be unique in a physical table
	Label string
	// DataType is the data type of the field
	DataType enum.DataType
	// Description is the description of the field
	Description string
	// CreatedAt is the created time of the field
	CreatedAt time.Time
	// UpdatedAt is the updated time of the field
	UpdatedAt time.Time
}
