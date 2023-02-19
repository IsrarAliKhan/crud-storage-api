package db

import (
	"crud-storage-api/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var connObj *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s TimeZone=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbName, config.DbPass, config.DbTz, config.DbSslMode)

	if config.DbSslMode != "disable" {
		dsn += fmt.Sprintf(" sslcert=%s sslkey=%s", config.DbSslCertFile, config.DbSslKeyFile)
	}

	connObj, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		FullSaveAssociations: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			DatabaseLocation, tzErr := time.LoadLocation(config.DbTz)
			if tzErr != nil {
				log.Fatalln("Cannot load database timezone info from .env, defaulting to UTC...")
				return time.Now().UTC().Truncate(time.Microsecond)
			}
			return time.Now().In(DatabaseLocation).Truncate(time.Microsecond)
		},
		PrepareStmt: true,
	})

	if err := Ping(); err != nil {
		log.Fatalln(err)
		return
	} else {
		log.Println("Successfully connected to database..")
	}
}

// Conn is globally accessible database handle
func Conn() *gorm.DB {
	if Ping() != nil {
		log.Panicln("Cannot connect to database..")
	}

	return connObj
}

// Ping will return error if database connection was unsuccessful
func Ping() error {
	db, err := connObj.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}

// PreloadUnscoped can be used as a callback function
// for Preload() to preload soft deleted records
func PreloadUnscoped(db *gorm.DB) *gorm.DB {
	return db.Unscoped()
}
