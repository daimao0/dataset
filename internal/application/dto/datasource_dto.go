package dto

import (
	"dataset/internal/common/enum"
	"time"
)

type DataSourceDTO interface {
	DataSourceDTOType() enum.DataSourceType
}

// BaseDataSourceDTO is the data transfer object for DataSource.
type BaseDataSourceDTO struct {
	// Id is the unique identifier of DataSource.
	Id int64
	// Name is the name of DataSource.
	Name string
	// Description is the description of DataSource.
	Description string
	// CreatedAt is the creation time of DataSource.
	CreatedAt time.Time
	// UpdatedAt is the last update time of DataSource.
	UpdatedAt time.Time
}
