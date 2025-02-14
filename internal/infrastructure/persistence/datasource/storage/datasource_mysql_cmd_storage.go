package storage

import (
	"dataset/internal/domain/datasource/model"
	"dataset/internal/infrastructure/persistence/db"
)

type DataSourceMySQLCmdStorage struct {
}

func NewDataSourceMySQLCmdStorage() *DataSourceMySQLCmdStorage {
	return &DataSourceMySQLCmdStorage{}
}

// Create  data source
func (d *DataSourceMySQLCmdStorage) Create(dataSource model.DataSource) error {
	db := db.GetDB()
	tx := db.Begin()
	tx.Create()
	return nil
}

// Update data source
func (d *DataSourceMySQLCmdStorage) Update(dataSource model.DataSource) error {

	return nil
}

// Delete data source
func (d *DataSourceMySQLCmdStorage) Delete(id string) error {

	return nil
}
