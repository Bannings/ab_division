package db

import (
	"ab_division/golbal"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var (
	dbABExperiment   *gorm.DB
	abExperimentOnce sync.Once
)

func OpenMysql(dbConfig golbal.DBConfig) (*gorm.DB, error) {

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&allowNativePasswords=true",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Schema)

	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(5)
	return db, nil
}

func ABExperimentDB() *gorm.DB {
	abExperimentOnce.Do(func() {
		conf := golbal.GetConfig()
		userConf := golbal.DBConfig{Schema: conf.ABMysql.Schema, Username: conf.ABMysql.Username, Password: conf.ABMysql.Password, Host: conf.ABMysql.Host, Port: conf.ABMysql.Port}
		db, err := OpenMysql(userConf)
		if err != nil {
			panic(err)
		}
		dbABExperiment = db
	})
	return dbABExperiment
}
