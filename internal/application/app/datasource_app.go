package app

import (
	"dataset/internal/application/cmd"
	"dataset/internal/application/dto"
)

// DataSourceApp The interface has many implementations for enum.DataSourceType
type DataSourceApp interface {

	// Create a data source by cmd.DataSourceCmd interface
	Create(cmd cmd.DataSourceCmd) (dto.DataSourceDTO, error)
}
