package handler

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/db"
)

type ApiPermission struct {
	Name          string `json:"name"`
	NamespacePerm string `json:"namespace_perm"`
	ServicePerm   string `json:"service_perm"`
	UserPerm      string `json:"user_perm"`
	AppPerm       string `json:"app_perm"`
}

type PermissionApiService struct{}

func (p *PermissionApiService) CreatePermission(c *gin.Context) {

	permission := db.Permission{}

	// check bind struct
	if err := c.ShouldBind(&permission); err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// setting created time time now
	permission.CreatedAt = time.Now()
	permission.UpdatedAt = time.Now()

	//
	if err := db.DB.Create(&permission).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{"code": http.StatusCreated, "data": permission, "msg": "创建成功!"})
}

func (p *PermissionApiService) DeletePermission(c *gin.Context) {

	//
	permission := db.Permission{}
	if err := c.ShouldBindUri(&permission); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// find model instance if exists
	if err := db.DB.Find(&permission).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// set is_delete true
	permission.IsDelete = true
	if err := db.DB.Save(&permission).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent,
		gin.H{"code": http.StatusNoContent, "data": "", "msg": "删除成功!"})
}

func (p *PermissionApiService) PermissionList(c *gin.Context) {

	var permissions []*db.Permission

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
	}, &permissions)

	// return
	c.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "data": paginator, "msg": ""})
}

func (p *PermissionApiService) EditPermission(c *gin.Context) {

	pid, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return
	}
	permission := db.Permission{Model: gorm.Model{ID: uint(pid)}}

	//fmt.Println(config)
	//permission.Value = permission.PermissionAuthorizeValue(config)

	if err := c.Bind(&permission); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// set updated time now
	permission.UpdatedAt = time.Now()

	if err := db.DB.Update(&permission).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "data": permission, "msg": "编辑成功"})
}

func (p *PermissionApiService) DetailPermission(c *gin.Context) {

	permission := db.Permission{}
	if err := c.ShouldBindUri(&permission); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// perm struct to map
	perMap := structs.Map(permission)
	perMap["config"] = permission.PermissionTransfer()

	c.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "data": perMap, "msg": ""})

}
