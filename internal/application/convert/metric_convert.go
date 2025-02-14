package convert

import (
	"dataset/internal/application/dto"
	"dataset/internal/domain/dataset/model"
)

func MetricToMetricDTO(metric *model.Metric) *dto.MetricDTO {
	return &dto.MetricDTO{
		Id:         metric.Id,
		Name:       metric.Name,
		Label:      metric.Label,
		DataType:   metric.DataType,
		CreateTime: metric.CreatedAt,
		UpdateTime: metric.UpdatedAt,
	}
}
