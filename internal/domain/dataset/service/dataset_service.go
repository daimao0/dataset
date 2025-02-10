package service

import (
	"dataset/internal/domain/dataset/args"
	"dataset/internal/domain/dataset/model"
)

// DatasetService is the interface of dataset application
type DatasetService interface {

	// GenerateDatasetByData creates a dataset by data
	GenerateDatasetByData(args *args.DatasetGenerateByDataArgs) (*model.Dataset, error)
}
