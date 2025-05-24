package dbConn

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, *sql.DB, error) {
	log.Println("Dbconnection (+) ")

	var lConstring string
	var lDb *sql.DB
	var lErr error
	var lGormDb *gorm.DB

	lDBCred := ReadDBCred()

	lConstring = `` + lDBCred.Conn1.User + `:` + lDBCred.Conn1.Password + `@tcp(` + lDBCred.Conn1.Server + `:` +
		fmt.Sprint(lDBCred.Conn1.Port) + `)/` + lDBCred.Conn1.Database + `?charset=utf8mb4&parseTime=True&loc=Local`
	lGormDb, lErr = gorm.Open(mysql.Open(lConstring), &gorm.Config{})
	if lErr != nil {
		log.Println("Error : ")
		return lGormDb, lDb, lErr
	}

	lDb, lErr = lGormDb.DB()
	if lErr != nil {
		log.Println("Error : ")
		return lGormDb, lDb, lErr
	}

	lDBConConfig := ConConfig()

	lDb.SetMaxIdleConns(lDBConConfig.OpenConnCt)

	lDb.SetMaxOpenConns(lDBConConfig.IdleConnCt)

	lDb.SetConnMaxIdleTime(time.Second * time.Duration(lDBConConfig.MaxIdleCt))

	log.Println("Dbconnection (-) ")
	return lGormDb, lDb, lErr
}

func GlobalDBConnection() error {
	var lErr error
	log.Println("GlobalDBConnection(+)")

	GConnInst.Conn1_gorm, GConnInst.Conn1_sql, lErr = GetDBConnection()
	log.Println("GlobalDBConnection(-)")
	return lErr

}
