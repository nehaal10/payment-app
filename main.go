package main

import (
	"errors"
	"flag"
	"fmt"
	"payment/app/internal/logger"
	"payment/app/setup"
	"payment/app/utils"

	"go.uber.org/zap/zapcore"
)

func parseCommandLineArgs() (string, string, string) {
	var dbupgrade string
	var dbInstall string
	var env string

	flag.StringVar(&dbupgrade, "dbupgrade", "", "fo dbupgrade")
	flag.StringVar(&dbInstall, "dbinstall", "", "for db installation")
	flag.StringVar(&env, "environment", "dev", "which environtment")

	flag.Parse()

	return dbupgrade, dbInstall, env
}

func main() {
	dbupgrade, dbInstall, env := parseCommandLineArgs()

	// inititate logger

	l := logger.NewCustomLogger(zapcore.DebugLevel)
	fmt.Println(l.Logger.Level())
	log := logger.GetLogger(l.Logger)
	log.Debug("sds", "ff", "dfsf", "Ff", "sfsf")
	log.Info("sds", "ff", "dfsf", "Ff", "sfsf")
	log.Error("sds", "ff", "dfsf", "Ff", "sfsf", errors.New("fdfs"))
	log.Fatal("sds", "ff", "dfsf", "Ff", "sfsf", errors.New("fdfs"))

	if len(dbInstall) > 0 {
		//install the DB
	}

	if len(dbupgrade) > 0 {
		// call the dbUpgrade script
	}
	if env != "dev" && env != "staging" {
		panic("env not there in the list")
	}
	config, err := utils.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	setup.Start(config.Server.Address, l.Logger)

}
