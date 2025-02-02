// CBSD Project 2013-2025
// beanstalkd driven CBSD sample
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	configFile	= flag.String("config", "config.json", "Path to config.json")
	broker		= flag.String("broker", "beanstalkd", "broker provider: 'beanstald' or 'rabbitmq'")
	cbsdEnv		= flag.String("cbsdenv", "/usr/jails", "CBSD workdir environment")
)

func check_cbsd_env (cbsdenv string) bool {

	name := fmt.Sprintf("%s/nc.inventory",cbsdenv)

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			fmt.Println(err.Error())
			return false
		}
	}

	return true
}

func main() {
	flag.Parse()

	config, err := LoadConfiguration(*configFile)

	if err != nil {
		fmt.Println("main: config load error")
		os.Exit(1)
	}

	log_init(config.Logfile)

	if len(config.CbsdEnv) > 0 {
		*cbsdEnv=config.CbsdEnv
	}
	if len(config.Broker) > 0 {
		*broker=config.Broker
	}

	if !check_cbsd_env(*cbsdEnv) {
		Fatal("CBSD env check error")
	}

	if config.CbsdColor == false {
		os.Setenv("NOCOLOR", "1")
	}

	Infof("Using config file: %s\n", *configFile)
	Infof("CBSD Env: %s\n", *cbsdEnv)
	Infof("Broker engine: %s\n", *broker)
	Infof("Logfile: %s\n", config.Logfile)
	Infof("MQ logdir: %s\n", config.BeanstalkConfig.LogDir)
	os.Setenv("cbsd_workdir", *cbsdEnv)
	os.Setenv("workdir", *cbsdEnv)
	os.Setenv("NOINTER", "1")

	beanstalkdLoop(config.BeanstalkConfig)
}
