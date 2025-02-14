package convert

import (
	"dataset/internal/adapter/http/request"
	"dataset/internal/application/cmd"
)

func MySQLDataSourceRequestToMySQLDataSourceCmd(request *request.MySQLDataSourceRequest) *cmd.MySQLDataSourceCmd {
	return &cmd.MySQLDataSourceCmd{
		BaseDataSourceCmd: cmd.BaseDataSourceCmd{
			Description: request.Description,
			Name:        request.Name,
		},
		Address:  request.Address,
		Database: request.Database,
		Password: request.Password,
		Username: request.Username,
	}
}
