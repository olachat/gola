package main

import (
	"os"

	"github.com/olachat/gola/golalib"
	"github.com/olachat/gola/mysqldriver"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("gola")
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	configPaths := []string{wd}
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	viper.ReadInConfig()
	viper.AutomaticEnv()
	driverName := "mysql"

	var config mysqldriver.Config = viper.GetStringMap(driverName)
	dbconfig := mysqldriver.NewDBConfig(config)
	output := config.DefaultString("output", "temp")
	gentype := config.DefaultString("gentype", "orm")

	golalib.Run(dbconfig, output, gentype)
}
