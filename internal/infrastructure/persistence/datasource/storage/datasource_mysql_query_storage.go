package storage

import "dataset/internal/domain/datasource/model"

type DataSourceMySQLQueryStorage struct {
}

func NewDataSourceMySQLQueryStorage() *DataSourceMySQLQueryStorage {
	return &DataSourceMySQLQueryStorage{}
}

// GetById get data source by id
func (d *DataSourceMySQLQueryStorage) GetById(id string) (model.DataSource, error) {
	return nil, nil
}

// ListAll list all data sources
func (d *DataSourceMySQLQueryStorage) ListAll() ([]model.DataSource, error) {

	return nil, nil
}
