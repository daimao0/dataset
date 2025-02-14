package convert

import (
	"dataset/internal/application/dto"
	"dataset/internal/domain/dataset/model"
)

func DatasetToDatasetDTO(dataset *model.Dataset) *dto.DatasetDTO {

	datasetDTO := &dto.DatasetDTO{
		ID:                dataset.Id,
		Name:              dataset.Name,
		Description:       dataset.Description,
		CreatedTime:       dataset.CreatedAt,
		UpdatedTime:       dataset.UpdatedAt,
		Measures:          make([]dto.MeasureDTO, 0),
		Metrics:           make([]dto.MetricDTO, 0),
		Dimensions:        make([]dto.DimensionDTO, 0),
		DerivedDimensions: make([]dto.DriverDimensionDTO, 0),
	}
	for _, field := range dataset.Fields {
		switch f := field.(type) {
		case *model.Measure:
			measureDTO := MeasureToMeasureDTO(f)
			datasetDTO.Measures = append(datasetDTO.Measures, *measureDTO)
			break
		case *model.Metric:
			metricDTO := MetricToMetricDTO(f)
			datasetDTO.Metrics = append(datasetDTO.Metrics, *metricDTO)
			break
		case *model.Dimension:
			dimensionDTO := DimensionToDimensionDTO(f)
			datasetDTO.Dimensions = append(datasetDTO.Dimensions, *dimensionDTO)
			break
		case *model.DriverDimension:
			driverDimensionDTO := DriverDimensionToDriverDimensionDTO(f)
			datasetDTO.DerivedDimensions = append(datasetDTO.DerivedDimensions, *driverDimensionDTO)
			break
		}
	}
	return datasetDTO
}
