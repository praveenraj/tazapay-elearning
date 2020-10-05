package initializer

import (
	"math/rand"
	"time"

	"tazapay.com/elearning/common/constants"
	"tazapay.com/elearning/common/driver"
)

// AppSetup application set up to be done for all invocations
// like establish database connection, any service client creation, dependency initiate, etc.
func AppSetup() (err error) {
	// seed the rand package before using it
	rand.Seed(time.Now().UnixNano())

	// init DB connection
	return initDatabase()
}

// initDatabase initialize the database connection.
func initDatabase() error {
	// DB connection input (Should be read from env variables)
	input := driver.Input{
		Driver:       driver.Mysql,
		SQLOrm:       driver.GOrm,
		Username:     "BGyfWKhCLq",
		Password:     "EMAy5lojpm",
		Host:         "remotemysql.com",
		Port:         3306,
		DatabaseName: "BGyfWKhCLq",
		ConnMaxTry:   constants.IntThree,
		Params: map[string]string{
			driver.InterPolateParams: constants.StrTrue,
			driver.CharSet:           "utf8mb4",
			driver.ParseTime:         constants.StrTrue,
			driver.Timeout:           "5s",
		},
	}
	// make database connection
	return driver.ConnectGorm(&input)
}
