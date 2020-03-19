package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/yametech/fuxi/pkg/db"
	"github.com/yametech/fuxi/thirdparty/lib/token"
)

type UserApiService struct {
	jwt  *token.Token
	pub  micro.Publisher
	wrap client.Wrapper
}

func NewUserApiService(pub micro.Publisher, token *token.Token, cliWrap client.Wrapper) *UserApiService {
	return &UserApiService{
		jwt:  token,
		pub:  pub,
		wrap: cliWrap,
	}
}

func (u *UserApiService) UserInfo(c *gin.Context) {
	// jwt rewrite header
	tokenUser := c.Request.Header.Get("x-auth-username")
	if len(tokenUser) < 1 {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": "验证有误"})
		return
	}
	user := &db.User{
		Name: &tokenUser,
	}

	if err := db.DB.Find(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}
	var m map[string]interface{}
	if err = json.Unmarshal(bytes, &m); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}

func (u *UserApiService) UserAuthorization(c *gin.Context) {

	user := &db.User{}
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	if err := db.DB.Find(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	expireTime := time.Now().Add(time.Minute * 30).Unix()
	s, err := u.jwt.Encode("", *user.Name, expireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// data
	data := map[string]interface{}{"user": user.Name, "id": user.ID, "token": s}
	c.JSON(http.StatusOK,
		gin.H{"code": http.StatusOK, "data": data, "msg": "编辑成功"})

}

func (u *UserApiService) UserRegister(c *gin.Context) {
	//Call lower layer service context TODO + ctx
	//ctx, ok := gin2micro.ContextWithSpan(c)
	//if ok == false {
	//	log.Error("get context err")
	//}
	//_ = ctx

	user := db.User{}

	// check bind struct
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// setting created time time now
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	//	create
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{"code": http.StatusCreated, "data": user, "msg": "创建成功!"})
}

func (u *UserApiService) UserDelete(c *gin.Context) {
	user := db.User{}
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// find model instance if exists
	if err := db.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	// set is_delete true
	user.IsDelete = true
	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusUnprocessableEntity, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent,
		gin.H{"code": http.StatusNoContent, "data": "", "msg": "删除成功!"})
}
