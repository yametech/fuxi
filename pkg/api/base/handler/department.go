package handler

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/db"
	"net/http"
	"strconv"
	"time"
)

type DepartmentApiService struct{}

func (d *DepartmentApiService) CreateDepartment(c *gin.Context) {

	department := db.Department{}

	// check bind struct
	if err := c.ShouldBind(&department); err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// setting created time time now
	department.CreatedAt = time.Now()
	department.UpdatedAt = time.Now()

	//	create
	if err := db.DB.Create(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{"code": http.StatusCreated, "data": department, "msg": "success!"})
}

func (d *DepartmentApiService) DeleteDepartment(c *gin.Context) {

	//
	department := db.Department{}
	if err := c.ShouldBindUri(&department); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// find model instance if exists
	if err := db.DB.Find(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// set is_delete true
	department.IsDelete = true
	if err := db.DB.Save(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent,
		gin.H{"code": http.StatusNoContent, "data": "", "msg": "deleted!"})
}

func (d *DepartmentApiService) DepartmentList(c *gin.Context) {

	var department []*db.Department

	// page
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	//query paginator
	paginator := pagination.Paging(&pagination.Param{
		DB:      db.DB.DB,
		Page:    page,
		Limit:   pageSize,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &department)

	// return
	c.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "data": paginator, "msg": ""})
}

func (d *DepartmentApiService) EditDepartment(c *gin.Context) {

	department := db.Department{}
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// set updated time now
	department.UpdatedAt = time.Now()

	if err := db.DB.Update(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "data": "", "msg": "success!"})
}
