package convert

import (
	"dataset/internal/common/easytool/convert"
	"dataset/internal/domain/datasource/model"
	"dataset/internal/infrastructure/persistence/dataset_db"
	"dataset/internal/infrastructure/persistence/datasource/po"
)

func DataSourceMySQLToDataSourcePO(dataSourceMySQL *model.DataSourceMySQL) *po.DataSourceMySQLPO {

	return &po.DataSourceMySQLPO{
		Model: db.Model{
			ID:        convert.ToInt64(dataSourceMySQL.Id),
			CreatedAt: dataSourceMySQL.CreateAt,
		},
		Address:  dataSourceMySQL.Address,
		Database: dataSourceMySQL.Database,
		Username: dataSourceMySQL.Username,
		Password: dataSourceMySQL.Password,
	}
}

func DataSourceMySQLToDataSource(dataSourceMySQL *model.DataSourceMySQL) *po.DataSourcePO {

	return &po.DataSourcePO{
		Model: db.Model{
			ID:        convert.ToInt64(dataSourceMySQL.Id),
			CreatedAt: dataSourceMySQL.CreateAt,
			UpdatedAt: dataSourceMySQL.UpdateAt,
		},
		Name:        dataSourceMySQL.Name,
		Description: dataSourceMySQL.Description,
	}
}
