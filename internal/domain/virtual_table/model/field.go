package model

import (
	"dataset/internal/common/enum"
	"time"
)

// Field is a field of a virtual table
type Field struct {
	// Id is the id of the field
	Id int64
	// Name is the name of the field, it must be unique in a virtual table
	// it Generates by the system,usually use to build sql
	Name string
	// Label is the label of the field,it must be unique in a virtual table
	Label string
	// Type is the type of the field
	Type enum.DataType
	// Description is the description of the field
	Description string
	// CreatedAt is the created time of the field
	CreatedAt time.Time
	// UpdatedAt is the updated time of the field
	UpdatedAt time.Time
}
