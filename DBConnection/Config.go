package dbConn

import (
	"database/sql"
	"fmt"
	"log"

	toml "Sales/common/toml"
	"strconv"

	"gorm.io/gorm"
)

type DBConnectionStruct struct {
	User     string
	Database string
	DBType   string
	Port     int
	Server   string
	DB       string
	Password string
}

type ConnInstStruct struct {
	Conn1_sql  *sql.DB
	Conn1_gorm *gorm.DB
}

type Db_DetailsStruct struct {
	Conn1 DBConnectionStruct
}

var GConnInst ConnInstStruct

// Struct for to hold the connection pool configuration
type ConConfigStruct struct {
	OpenConnCt int
	IdleConnCt int
	MaxIdleCt  int
}

/*
Method will read the database detail from the toml
Ex : Userdetail : "root", port: 3306 etc..,
*/
func ReadDBCred() Db_DetailsStruct {
	log.Println("ReadDBCred (+) ")

	var lDBCred Db_DetailsStruct

	lConfig := toml.ReadTomlFile("./toml/dbconfig.toml")
	lDBCred.Conn1.User = fmt.Sprintf("%v", lConfig.(map[string]interface{})["DBUser"])
	lDBCred.Conn1.Server = fmt.Sprintf("%v", lConfig.(map[string]interface{})["DBServer"])
	lDBCred.Conn1.Password = fmt.Sprintf("%v", lConfig.(map[string]interface{})["DBPassword"])
	lDBCred.Conn1.Database = fmt.Sprintf("%v", lConfig.(map[string]interface{})["DBDatabase"])
	lDBCred.Conn1.Port, _ = strconv.Atoi(fmt.Sprintf("%v", lConfig.(map[string]interface{})["DBPort"]))
	lDBCred.Conn1.DB = fmt.Sprintf("%v", lConfig.(map[string]interface{})["DBName"])
	lDBCred.Conn1.DBType = fmt.Sprintf("%v", lConfig.(map[string]interface{})["DBType"])

	log.Println("ReadDBCred (-) ")
	return lDBCred
}

func ConConfig() ConConfigStruct {
	log.Println("ConConfig (+) ")

	var ConnRes ConConfigStruct
	var lErr error

	// reading a connection pool details from the toml
	lDbConnectionpool := toml.ReadTomlFile("./toml/dbconfig.toml")
	lSetMaxOpenConns := fmt.Sprintf("%v", lDbConnectionpool.(map[string]interface{})["SetMaxOpenConnsdb"])
	lSetMaxIdleConnsdb := fmt.Sprintf("%v", lDbConnectionpool.(map[string]interface{})["SetMaxIdleConnsdb"])
	lSetConnMaxIdleTime := fmt.Sprintf("%v", lDbConnectionpool.(map[string]interface{})["SetConnMaxIdleTimedb"])

	// If the details not properly readen from the toml file this will handle the issue
	// if lSetMaxOpenConns == "" {
	// 	lSetMaxOpenConns = "10"
	// }

	// if lSetMaxIdleConnsdb == "" {
	// 	lSetMaxIdleConnsdb = "5"
	// }

	// if lSetConnMaxIdleTime == "" {
	// 	lSetConnMaxIdleTime = "5"
	// }

	ConnRes.OpenConnCt, lErr = strconv.Atoi(lSetMaxOpenConns)
	if lErr != nil {
		log.Println("Error :")
		return ConnRes
	}

	ConnRes.IdleConnCt, lErr = strconv.Atoi(lSetMaxIdleConnsdb)
	if lErr != nil {
		log.Println("Error :")
		return ConnRes
	}

	ConnRes.MaxIdleCt, lErr = strconv.Atoi(lSetConnMaxIdleTime)
	if lErr != nil {
		log.Println("Error :")
		return ConnRes
	}

	log.Println("ConConfig (-) ")
	return ConnRes
}
