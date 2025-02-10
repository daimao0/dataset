package impl

import (
	"dataset/internal/application/cmd"
	"dataset/internal/application/dto"
	"dataset/internal/domain/dataset/args"
	"dataset/internal/domain/dataset/service"
	"dataset/internal/domain/dataset/service/impl"
	"dataset/internal/infrastructure/convert"
)

type DatasetAppImpl struct {
	datasetService service.DatasetService
}

func NewDatasetAppImpl() *DatasetAppImpl {
	return &DatasetAppImpl{datasetService: impl.NewDatasetServiceImpl()}
}

// CreateDatasetByData generates a dataset from user uploads data.
// @param datasetCreateCmd: the dataset creation command.
// @return dto.DatasetDTO : the dataset.
func (ds *DatasetAppImpl) CreateDatasetByData(appCmd *cmd.DatasetCreateByDataCmd) (*dto.DatasetDTO, error) {
	dataArgs := &args.DatasetGenerateByDataArgs{Name: appCmd.Name, Data: appCmd.Data}
	dataset, _ := ds.datasetService.GenerateDatasetByData(dataArgs)
	datasetDTO := convert.DatasetToDatasetDTO(dataset)
	return datasetDTO, nil
}
