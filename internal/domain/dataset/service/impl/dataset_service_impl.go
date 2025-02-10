package impl

import (
	"dataset/internal/common/easytool/util/id_util"
	"dataset/internal/domain/dataset/args"
	"dataset/internal/domain/dataset/model"
	"dataset/internal/domain/dataset/service"
	"time"
)

// DatasetServiceImpl handles the dataset domain service
type DatasetServiceImpl struct {
	fieldService service.FieldService
}

func NewDatasetServiceImpl() *DatasetServiceImpl {
	return &DatasetServiceImpl{
		fieldService: NewFieldServiceImpl(),
	}
}

// GenerateDatasetByData generate dataset by data
func (ds *DatasetServiceImpl) GenerateDatasetByData(args *args.DatasetGenerateByDataArgs) (*model.Dataset, error) {
	// generate dataset
	name := args.Name
	if args.Name == "" {
		name = generateDatasetName()
	}
	dataset := &model.Dataset{Id: id_util.GenID(), Name: name, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	// parse data to generate fields,merge map
	mergeData := mergeKey(args.Data)
	fields, _ := ds.fieldService.GenerateFieldByMap(mergeData)
	dataset.Fields = fields
	return dataset, nil
}

// mergeKey range the data to merge every rows key
func mergeKey(data []map[string]any) map[string]any {
	mergeData := make(map[string]any)
	for _, row := range data {
		for k, v := range row {
			mergeData[k] = v
		}
	}
	return mergeData
}

// generateDatasetName generate default dataset name
func generateDatasetName() string {
	uuid, _ := id_util.GenerateShortUUID(8)
	return "t_" + uuid
}
