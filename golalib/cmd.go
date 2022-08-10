package golalib

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/olachat/gola/mysqldriver"
	"github.com/spf13/viper"
)

type cmd struct {
}

func GenCommandFactory() (cli.Command, error) {
	return genCmd, nil
}

var genCmd *cmd

func (*cmd) Help() string {
	return "generate orm stubs from default toml config"
}

func (*cmd) Run(args []string) int {
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

	Run(dbconfig, output)

	return 0
}

func (*cmd) Synopsis() string {
	return "generate orm stubs from default toml config"
}
