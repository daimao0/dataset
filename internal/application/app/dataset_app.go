package app

import (
	"dataset/internal/application/cmd"
	"dataset/internal/application/dto"
)

// DatasetApplication is the interface of dataset application
type DatasetApplication interface {

	// CreateDatasetByData parse data to create dataset
	CreateDatasetByData(cmd *cmd.DatasetCreateByDataCmd) (*dto.DatasetDTO, error)
}
