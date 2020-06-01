package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Options func(*Option)

type Option struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewOption(opts ...Options) *Option {
	opt := &Option{
		Host:     "127.0.0.1",
		Port:     3306,
		DBName:   "default",
		User:     "default",
		Password: "",
	}
	for i := range opts {
		(opts[i])(opt)
	}
	return opt
}

// WithOptionsCreateClient use a one or more option create mysql clientv2
func WithOptionsCreateClient(mo *Option) (*Client, error) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			mo.User,
			mo.Password,
			fmt.Sprintf("%s:%d", mo.Host, mo.Port),
			mo.DBName,
		),
	)
	if err != nil {
		return nil, err
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	db.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return &Client{db}, nil
}

func SetOptionHost(host string) Options {
	return func(o *Option) {
		o.Host = host
	}
}

func SetOptionPort(port int) Options {
	return func(o *Option) {
		o.Port = port
	}
}

func SetOptionDB(db string) Options {
	return func(o *Option) {
		o.DBName = db
	}
}

func SetOptionUser(user string) Options {
	return func(o *Option) {
		o.User = user
	}
}

func SetOptionPassword(password string) Options {
	return func(o *Option) {
		o.Password = password
	}
}
