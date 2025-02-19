package impl

import (
	"dataset/internal/application/cmd"
	"dataset/internal/application/dto"
	"dataset/internal/application/factory"
	"dataset/internal/domain/datasource/repository"
	"fmt"
)

// DataSourceMySQLAppImpl implements app.DataSourceApp
type DataSourceMySQLAppImpl struct {
	dataSourceMap map[repository.DataSourceCmdRepository]repository.DataSourceCmdRepository
}

// NewDataSourceMySQLAppImpl creates a new DataSourceMySQLAppImpl
func NewDataSourceMySQLAppImpl() *DataSourceMySQLAppImpl {
	return &DataSourceMySQLAppImpl{}
}

// Create a mysql datasource
func (d *DataSourceMySQLAppImpl) Create(dataSourceCmd cmd.DataSourceCmd) (dto.DataSourceDTO, error) {
	// generate datasource
	dataSource := factory.CreateDataSource(dataSourceCmd)
	//
	fmt.Println(dataSource)
	return nil, nil
}
