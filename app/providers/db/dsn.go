package db

import "fmt"

type Dsn struct {
	Driver string
}

type MysqlDsn struct {
	*Dsn
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Charset  string
}

func NewMySqlDsn(dsn *MysqlDsn) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dsn.Username, dsn.Password,
		dsn.Host, dsn.Port, dsn.Database,
		dsn.Charset,
	)
}
