package factory

import (
	"dataset/internal/application/cmd"
	"dataset/internal/domain/datasource/model"
)

func CreateDataSource(dataSourceCmd cmd.DataSourceCmd) model.DataSource {
	switch obj := dataSourceCmd.(type) {
	case *cmd.MySQLDataSourceCmd:
		return &model.DataSourceMySQL{
			BaseDataSource: model.BaseDataSource{Id: obj.Id, Name: obj.Name, Description: obj.Description},
			Address:        obj.Address,
			Database:       obj.Database,
			Password:       obj.Password,
			Username:       obj.Username,
		}
	}
	return nil
}
