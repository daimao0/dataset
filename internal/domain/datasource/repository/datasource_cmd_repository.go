package repository

import (
	"dataset/internal/domain/datasource/model"
)

// DataSourceCmdRepository by CQRS
type DataSourceCmdRepository interface {

	// Create data source
	Create(dataSource model.DataSource) error

	// Update data source
	Update(dataSourceCmd model.DataSource) error

	// Delete data source
	Delete(id string) error
}
