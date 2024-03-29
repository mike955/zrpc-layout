package dao

import (
	"fmt"
	"os"
	"time"

	configs "github.com/mike955/zrpc-layout/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type CommonTimeModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Init(config configs.Mysql) {
	var err error

	addr, username, password, database := config.MysqlAddr, config.MysqlUsername, config.MysqlPassword, config.MysqlDatabase

	if os.Getenv("MYSQL_ADDR") != "" && os.Getenv("MYSQL_PORT") != "" {
		addr = os.Getenv("MYSQL_ADDR") + ":" + os.Getenv("MYSQL_PORT")
	}
	if os.Getenv("MYSQL_USERNAME") != "" {
		username = os.Getenv("MYSQL_USERNAME")
	}
	if os.Getenv("MYSQL_PASSWORD") != "" {
		password = os.Getenv("MYSQL_PASSWORD")
	}
	if os.Getenv("MYSQL_DATABASE") != "" {
		database = os.Getenv("MYSQL_DATABASE")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, database)
	DB, err = gorm.Open(mysql.Open(dsn))
	DB = DB.Debug()
	if err != nil {
		panic("connect mysql error: " + err.Error())
	}
}
