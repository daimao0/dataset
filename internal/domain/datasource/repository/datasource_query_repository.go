package repository

import (
	"dataset/internal/domain/datasource/model"
)

// DataSourceQueryRepository by CQRS
type DataSourceQueryRepository interface {

	// GetById get data source by id
	GetById(id string) (model.DataSource, error)

	// ListAll list all data sources
	ListAll() ([]model.DataSource, error)
}
