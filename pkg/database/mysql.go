package database

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	mysqlConfig "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"shortlink-system/pkg/entities"
	"shortlink-system/pkg/helper"
	"time"
)

func InitializedDatabase() *gorm.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", username, password, host, port, dbName)
	if os.Getenv("DB_IS_SECURE") == "true" {
		caCert, err := os.ReadFile(fmt.Sprintf("cert-mysql/%s", os.Getenv("DB_CA")))
		helper.IfErrorHandler(err)
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			panic("Failed to append CA cert")
		}
		tlsConfig := &tls.Config{
			RootCAs: caCertPool,
		}
		mysqlConfig.RegisterTLSConfig("custom", tlsConfig)
		dsn += "&tls=custom"
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	helper.IfErrorHandler(err)
	log.Info("Connected to database")

	dbSql, err := db.DB()
	helper.IfErrorHandler(err)

	dbSql.SetMaxIdleConns(10)
	dbSql.SetMaxOpenConns(100)
	dbSql.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.ShortedLink{})

	return db
}
