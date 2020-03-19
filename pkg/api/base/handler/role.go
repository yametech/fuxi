package handler

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/db"
)

type RoleApiService struct{}

func (r *RoleApiService) CreateRole(c *gin.Context) {

	role := db.Role{}

	// check bind struct
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// setting created time time now
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	//	create
	if err := db.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{"code": http.StatusCreated, "data": role, "msg": "创建成功!"})
}

func (r *RoleApiService) DeleteRole(c *gin.Context) {

	//
	role := db.Role{}
	if err := c.ShouldBindUri(&role); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// find model instance if exists
	if err := db.DB.Find(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// set is_delete true
	role.IsDelete = true
	if err := db.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent,
		gin.H{"code": http.StatusNoContent, "data": "", "msg": "删除成功!"})
}

func (r *RoleApiService) RoleList(c *gin.Context) {

	var roles []*db.Role

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
	}, &roles)

	// return
	c.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "data": paginator, "msg": ""})
}

func (r *RoleApiService) EditRole(c *gin.Context) {

	role := db.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// set updated time now
	role.UpdatedAt = time.Now()

	if err := db.DB.Update(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "data": role, "msg": "编辑成功"})
}
