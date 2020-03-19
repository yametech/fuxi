## Reference cmd/base (github.com/yametech/cmd/base/main.go)
```

import (

    // swagger doc
	file "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"

    // import you project docs
	_ "github.com/yametech/fuxi/cmd/base/docs"
)

// title head 

// @title Gin swagger
// @version 1.0
// @description Gin swagger base
// @contact.name laik author
// @contact.url  github.com/yametech
// @contact.email laik.lj@me.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html


// User info doc
// @Summary base service user into
// @Description User info query
// @Tags user
// @Accept mpfd
// @Produce json
// @Param user query string true "user_id"
// @Success 200 {string} json "{"msg": "user_name"}"
// @Failure 400 {string} json "{"msg": "Please login"}"
// @Router /base/v1/user [get]
func Function(){...}

func main(){
    // ................

    // Then, if you set envioment variable NAME_OF_ENV_VARIABLE to anything, /swagger/*any will respond 404, just like when route unspecified.
    // Release production environment can be turned on
    router.GET("/base/swagger/*any", swag.DisablingWrapHandler(file.Handler, "NAME_OF_ENV_VARIABLE"))

}



make gen-openapi
```

## Official Reference book
https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html

## Parameters that separated by spaces. 
```
param name,param type,data type,is mandatory?,comment attribute(optional)
```

## Param Type
```
object (struct)
string (string)
integer (int, uint, uint32, uint64)
number (float32)
boolean (bool)
array
```

## Data Type
```
string (string)
integer (int, uint, uint32, uint64)
number (float32)
boolean (bool)
user defined struct
```

## Attribute
```
// @Param enumstring query string false "string enums" Enums(A, B, C)
// @Param enumint query int false "int enums" Enums(1, 2, 3)
// @Param enumnumber query number false "int enums" Enums(1.1, 1.2, 1.3)
// @Param string query string false "string valid" minlength(5) maxlength(10)
// @Param int query int false "int valid" mininum(1) maxinum(10)
// @Param default query string false "string default" default(A)
```

## Match pattern
```
path parameters, such as /users/{id}
query parameters, such as /users?role=admin
header parameters, such as X-MyHeader: Value
cookie parameters, which are passed in the Cookie header, such as Cookie: debug=0; csrftoken=BUSe35dohU3O1MZvDCU
```



