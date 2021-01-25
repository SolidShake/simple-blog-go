package db

import (
	"github.com/SolidShake/simple-blog-go/internal/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func InitDB(cnf config.Config) {
	var err error
	DB, err = sqlx.Connect("mysql", getDataSourceName(cnf))
	if err != nil {
		panic(err)
	}
}

func getDataSourceName(cnf config.Config) string {
	dbConf := cnf.Database
	return dbConf.Username + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/" + dbConf.DbName + "?parseTime=true"
}
