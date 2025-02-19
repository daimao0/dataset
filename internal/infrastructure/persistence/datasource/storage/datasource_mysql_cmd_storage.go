package storage

import (
	"dataset/internal/domain/datasource/model"
	"dataset/internal/infrastructure/convert"
	"dataset/internal/infrastructure/persistence/dataset_db"
)

type DataSourceMySQLCmdStorage struct {
}

func NewDataSourceMySQLCmdStorage() *DataSourceMySQLCmdStorage {
	return &DataSourceMySQLCmdStorage{}
}

// Create  data source
func (d *DataSourceMySQLCmdStorage) Create(dataSource model.DataSource) error {
	dataSourceMySQL := dataSource.(*model.DataSourceMySQL)
	dataSourcePO := convert.DataSourceMySQLToDataSourcePO(dataSourceMySQL)
	dataSourceMySQLPO := convert.DataSourceMySQLToDataSourcePO(dataSourceMySQL)
	db := dataset_db.GetDB()
	tx := db.Begin()
	tx.Create(dataSourcePO)
	tx.Create(dataSourceMySQLPO)
	tx.Commit()
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
