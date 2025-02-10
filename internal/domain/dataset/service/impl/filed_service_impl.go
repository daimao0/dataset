package impl

import (
	"dataset/internal/common/easytool/util/id_util"
	"dataset/internal/common/enum"
	"dataset/internal/common/error_code"
	"dataset/internal/domain/dataset/model"
	"time"
)

type FieldServiceImpl struct {
}

func NewFieldServiceImpl() *FieldServiceImpl {
	return &FieldServiceImpl{}
}

// GenerateFieldByMap generate field by map
func (fs *FieldServiceImpl) GenerateFieldByMap(row map[string]any) ([]model.Field, error) {
	var fields []model.Field
	for k, v := range row {
		field, err := generateField(k, v)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}
	return fields, nil
}

// generateField generate field by key and value, key will parse to field metadata , value will parse to field data type
func generateField(key string, value any) (model.Field, error) {
	// if value is nil, then return dimension field
	if value == nil {
		dim := &model.Dimension{
			Id:        id_util.GenID(),
			Name:      genFieldName(),
			Label:     key,
			DataType:  enum.String,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		return dim, nil
	}
	// parse string value to data type
	dataType := enum.ParseToDataType(value)
	switch dataType {
	case enum.Number:
		return &model.Measure{
			Id:        id_util.GenID(),
			Name:      genFieldName(),
			Label:     key,
			DataType:  dataType,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	case enum.DateTime:
	case enum.Enum:
	case enum.String:
		return &model.Dimension{
			Id:        id_util.GenID(),
			Name:      genFieldName(),
			Label:     key,
			DataType:  dataType,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	default:
		return nil, error_code.FieldTypeNotExist
	}
	return nil, nil
}

// genFieldName generate field name
func genFieldName() string {
	shortUUID, _ := id_util.GenerateShortUUID(8)
	return "f_" + shortUUID
}
