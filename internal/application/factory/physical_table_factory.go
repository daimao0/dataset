package factory

import (
	"dataset/internal/application/cmd"
	"dataset/internal/common/easytool/util/id_util"
	"dataset/internal/common/easytool/util/slice_util"
	"dataset/internal/domain/physical_table/model"
	"time"
)

// CreatePhysicalTableFormData is used to create a physical table from data.
func CreatePhysicalTableFormData(cmd *cmd.PhysicalTableCreateByDataCmd) *model.PhysicalTable {

	// generate dataset
	name := cmd.Name
	if cmd.Name == "" {
		name = generateDatasetName()
	}
	// parse data to generate fields,merge map
	data := slice_util.RandomSampling(cmd.Data, 10)
	mergeData := mergeKey(data)
	fields, _ := GenerateFieldByMap(mergeData)
	return &model.PhysicalTable{
		Id:          id_util.GenID(),
		Name:        name,
		Fields:      fields,
		OriginName:  cmd.Name,
		Description: "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
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
