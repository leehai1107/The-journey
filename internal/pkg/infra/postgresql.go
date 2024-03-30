package infra

import (
	"fmt"
	"log"
	"time"

	"github.com/leehai1107/The-journey/internal/pkg/config"
	"github.com/leehai1107/The-journey/internal/pkg/constant"
	"github.com/leehai1107/The-journey/internal/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbSingleton *gorm.DB

func InitPostgresql() {
	dbCfg := config.DBConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		dbCfg.PgHost,
		dbCfg.PgUser,
		dbCfg.PgPassword,
		dbCfg.PgDatabase,
		dbCfg.PgPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		logger.Fatalf("%s %s", err.Error(), "unable to instantiate database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("%s %s", err.Error(), "unable to get sql.DB from gorm.DB")
	}

	if dbCfg.PgPoolSize > 0 {
		sqlDB.SetMaxOpenConns(dbCfg.PgPoolSize)
	}

	if dbCfg.PgIdleConnTimeout > 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(dbCfg.PgIdleConnTimeout) * time.Second)
	}

	if dbCfg.PgMaxConnAge > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.PgMaxConnAge) * time.Second)
	}

	if config.ServerConfig().ENV != constant.ProductionEnv {
		db = db.Debug()
	}

	dbSingleton = db
}

func ClosePostgresql() error {
	sqlDB, err := dbSingleton.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func GetDB() *gorm.DB {
	if dbSingleton == nil {
		log.Printf("Connection to Database Postgres is not setup")
	}

	return dbSingleton
}

// BeginTransaction start an Transaction, require defer ReleaseTransaction instantly
func BeginTransaction() *gorm.DB {
	return dbSingleton.Begin()
}

func ReleaseTransaction(tx *gorm.DB) {
	tx.Commit()
}
