package main

import (
	"fmt"
	"time"
	"flag"
	"github.com/matishsiao/goInfo"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Server			string
	Port			int
	Key			string
	Connection_String	string
	Hostname		string
}

func main() {
	fmt.Println("##################################################")
	//Getting Start time
	startdt := time. Now()
	fmt.Println("StartTime:", startdt)
	//Getting Commmands Line flags
	configPtr := flag.String("config", "", "Config File path")
	flag.Parse()
	configfile := *configPtr
	fmt.Println("Configfile:", configfile)

	//Get settings from Config File
	fmt.Println("####################Config Settings###############")
	configuration := Configuration{}
	err := gonfig.GetConf(configfile, &configuration)
	if err != nil {
		panic(err)
	}
	serverurl := configuration.Server
	serverport := configuration.Port
	serverkey := configuration.Key
	fmt.Println("ServerUrl:", serverurl)
	fmt.Println("ServerPort:", serverport)
	fmt.Println("ServerKey:", serverkey)
	fmt.Println("##################################################")

	//Getting Hostname
	if configuration.Hostname != "" {
		fmt.Println("Hostname override")
	}
	gi := goInfo.GetInfo()
	hostname := gi.Hostname
	fmt.Println("Hostname:", hostname)
}
