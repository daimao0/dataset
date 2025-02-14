package controller

import (
	"dataset/internal/adapter/http/request"
	"dataset/internal/application/app"
	"dataset/internal/application/app/impl"
	"dataset/internal/application/convert"
	"dataset/internal/common/easytool/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DataSourceController is a controller for data source
type DataSourceController struct {
	dataSourceApp app.DataSourceApp
}

// NewDataSourceController is a constructor for DataSourceController
func NewDataSourceController() *DataSourceController {
	return &DataSourceController{
		dataSourceApp: impl.NewDataSourceMySQLAppImpl(),
	}
}

// UploadFileToCreateDataSource is a method for DataSourceController
func (d *DataSourceController) UploadFileToCreateDataSource(c *gin.Context) {

}

// CreateMySQLDataSource use form to create data source
func (d *DataSourceController) CreateMySQLDataSource(c *gin.Context) {
	req := &request.MySQLDataSourceRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	mySQLDataSourceCmd := convert.MySQLDataSourceRequestToMySQLDataSourceCmd(req)
	create, err := d.dataSourceApp.Create(mySQLDataSourceCmd)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(create))
}
