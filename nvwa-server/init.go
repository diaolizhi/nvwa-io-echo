package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nvwa-io/base"
	"os"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func initDB() {
	var err error
	base.DB, err = gorm.Open(mysql.Open(viper.GetString("database.uri")), &gorm.Config{
		//Logger: logger,
	})

	if err != nil {
		//log.Fatal(errors.Wrap(err, "could not init database"))
		os.Exit(-1)
	}
}
