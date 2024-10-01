package main

import (
	"flag"
	"payment/app/setup"
	"payment/app/utils"
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

	setup.Start(config.Server.Address)

}
