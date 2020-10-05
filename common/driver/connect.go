package driver

import (
	"fmt"

	"tazapay.com/elearning/common/constants"

	_ "github.com/go-sql-driver/mysql" // mysql connectivity driver
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// global object
var (
	_db *DB
)

// GetDB return DB struct
func GetDB() *DB {
	return _db
}

// SetDB set DB struct
func SetDB(db *DB) {
	_db = db
}

// GetDriver return driver name being used
func (d *DB) GetDriver() string {
	return d.Driver
}

// SetDriver set driver name to be used
func (d *DB) SetDriver(name string) {
	d.Driver = name
}

// GetOrm return orm name being used
func (d *DB) GetOrm() string {
	return d.Orm
}

// SetOrm set orm name to be used
func (d *DB) SetOrm(name string) {
	d.Orm = name
}

// GetGOrm return gorm object
func (d *DB) GetGOrm() *gorm.DB {
	return d.GOrm
}

// SetGOrm set gorm object
func (d *DB) SetGOrm(db *gorm.DB) {
	d.GOrm = db
}

// ConnectGorm connect database with gorm.
// @param i: database connection inputs
func ConnectGorm(i *Input) error {
	var err error
	// Check if a connection is already open
	if GetDB() != nil && GetDB().GetGOrm() != nil {
		return nil
	}

	// get connection string
	connectionString := getConnectionString(i)
	log.Debug().Msgf("connection URI: %v", connectionString)

	// try to connect the database until maxtry count drains
	// in any of the successive attempt fails
	for currentTry := constants.IntOne; currentTry <= i.ConnMaxTry; currentTry++ {
		// register default database
		var conn *gorm.DB
		conn, err = gorm.Open(i.Driver, connectionString)
		if err == nil {
			// check mysql connection
			err = conn.DB().Ping()
			if err == nil {
				conn.LogMode(constants.BoolTrue)
				// set driver DB to use the conn & config at global level
				initGlobalDBObj(i, conn)
			}
		}
		break
	}
	return err
}

// getConnectionString frame the connection string with the input params
func getConnectionString(i *Input) string {
	connectionString := fmt.Sprintf(mysqlConnURIFormat, i.Username, i.Password, i.Host, i.Port, i.DatabaseName)

	if len(i.Params) > constants.IntZero {
		connectionString += constants.QuestionMark
	}

	// append the query params
	var paramPrefix string
	for k, v := range i.Params {
		connectionString += paramPrefix + k + constants.Equal + v
		paramPrefix = constants.Ampersand
	}

	return connectionString
}

// initGlobalDBObj set application centric values to DB obj and made it global
func initGlobalDBObj(i *Input, args ...interface{}) {
	// init new db obj
	db := new(DB)

	// set driver & orm name
	db.SetDriver(i.Driver)
	db.SetOrm(i.SQLOrm)

	for _, arg := range args {
		switch v := arg.(type) {
		case *gorm.DB:
			db.SetGOrm(v)

		case gorm.DB:
			db.SetGOrm(&v)
		}
	}

	// set the db obj to the global variable
	SetDB(db)
}

// ConnectMock mock database connection for unit testing
func ConnectMock() error {
	i := &Input{
		SQLOrm: Mock,
	}
	initGlobalDBObj(i)
	return nil
}
