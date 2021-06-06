package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := NewMySqlDsn(&MysqlDsn{
		Host:     "192.168.8.117",
		Port:     3306,
		Database: "t591_new",
		Username: "t591",
		Password: "sogamysql",
		Charset:  "utf8",
	})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
