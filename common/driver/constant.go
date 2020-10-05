package driver

import (
	"github.com/jinzhu/gorm"
)

// DB data management object
type DB struct {
	Driver string
	Orm    string
	GOrm   *gorm.DB
}

// Input required values to make database connection
type Input struct {
	Driver       string
	SQLOrm       string
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
	Params       map[string]string
	ConnMaxTry   int
}

// drivers & ORMs
const (
	Mysql = "mysql"
	GOrm  = "gorm"
	Mock  = "mock"
)

// connection URI
const (
	mysqlConnURIFormat = "%s:%s@tcp(%s:%d)/%s"
)

// connection params
const (
	InterPolateParams = "interpolateParams"
	CharSet           = "charset"
	ParseTime         = "parseTime"
	Timeout           = "timeout"
)
