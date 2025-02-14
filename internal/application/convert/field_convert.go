package convert

import (
	"dataset/internal/application/dto"
	"dataset/internal/domain/physical_table/model"
)

func FieldToFieldDTO(field *model.Field) *dto.FieldDTO {
	fieldDTO := &dto.FieldDTO{
		Id:          field.Id,
		Name:        field.Name,
		Label:       field.Label,
		DataType:    field.DataType,
		Description: field.Description,
		CreatedAt:   field.CreatedAt,
		UpdatedAt:   field.UpdatedAt,
	}
	return fieldDTO
}
func FieldSToFieldDTOS(fields []*model.Field) []*dto.FieldDTO {
	fieldDTOS := make([]*dto.FieldDTO, 0, len(fields))
	for _, field := range fields {
		fieldDTOS = append(fieldDTOS, FieldToFieldDTO(field))
	}

	return fieldDTOS
}
