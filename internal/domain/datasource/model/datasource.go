package model

import (
	"dataset/internal/common/enum"
	"time"
)

// DataSource domain
type DataSource interface {

	// DataSourceType The function is used to distinguish between different data source
	DataSourceType() enum.DataSourceType
}

// BaseDataSource is the base data source
type BaseDataSource struct {
	// Id is the id of the data source,It is a unique identifier
	Id string

	// Name is the name of the data source,It is created by the user
	Name string

	// Description is the description of the data source,It is created by the user
	Description string

	// CreateAt is the created time of the data source
	CreateAt time.Time

	// UpdateAt is the update time of the data source
	UpdateAt time.Time
}
