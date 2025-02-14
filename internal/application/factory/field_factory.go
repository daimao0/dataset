package factory

import (
	"dataset/internal/common/easytool/util/id_util"
	"dataset/internal/common/enum"
	"dataset/internal/domain/physical_table/model"
	"time"
)

// GenerateFieldByMap generate field by map
func GenerateFieldByMap(row map[string]any) ([]*model.Field, error) {
	var fields []*model.Field
	for k, v := range row {
		field := generateField(k, v)
		fields = append(fields, field)
	}
	return fields, nil
}

// generateField generate field by key and value, key will parse to field metadata , value will parse to field data type
func generateField(key string, value any) *model.Field {
	// if value is nil, then return dimension field
	field := &model.Field{
		Id:        id_util.GenID(),
		Name:      genFieldName(),
		Label:     key,
		DataType:  enum.ParseToDataType(value),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return field
}

// genFieldName generate field name
func genFieldName() string {
	shortUUID, _ := id_util.GenerateShortUUID(8)
	return "f_" + shortUUID
}
