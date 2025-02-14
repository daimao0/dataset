package impl

import (
	"dataset/internal/application/cmd"
	"dataset/internal/application/dto"
	"dataset/internal/application/factory"
	"fmt"
)

// DataSourceMySQLAppImpl implements app.DataSourceApp
type DataSourceMySQLAppImpl struct {
}

// NewDataSourceMySQLAppImpl creates a new DataSourceMySQLAppImpl
func NewDataSourceMySQLAppImpl() *DataSourceMySQLAppImpl {
	return &DataSourceMySQLAppImpl{}
}

// Create a mysql datasource
func (d *DataSourceMySQLAppImpl) Create(dataSourceCmd cmd.DataSourceCmd) (dto.DataSourceDTO, error) {
	dataSource := factory.CreateDataSource(dataSourceCmd)
	fmt.Println(dataSource)
	return nil, nil
}
