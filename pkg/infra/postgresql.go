package infra

import (
	"errors"
	"fmt"
	"ganja/pkg/server/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var postgresql *gorm.DB

func Postgresql() *gorm.DB {
	if postgresql == nil {
		panic(errors.New("setup postgresql first"))
	}
	return postgresql
}

func setupPostgresql() {
	conf := viper.GetViper()
	host := conf.GetString("postgresql.host")
	port := conf.GetInt("postgresql.port")
	user := conf.GetString("postgresql.user")
	password := conf.GetString("postgresql.password")
	db := conf.GetString("postgresql.db")
	connString := fmt.Sprintf("host=%v port=%d user=%v dbname=%v password=%v sslmode=disable", host, port, user, db, password)
	var err error
	postgresql, err = gorm.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	logrus.Info("Setup postgesql successfully")
	initModel()
}

func initModel() {
	Postgresql().AutoMigrate(&entity.Customer{})
}
