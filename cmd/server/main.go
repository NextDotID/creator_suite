package main

import (
	"flag"
	"fmt"
	"github.com/nextdotid/creator_suite/config"
	"github.com/nextdotid/creator_suite/controller"
	"github.com/nextdotid/creator_suite/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

const (
	ENVIRONMENT    = "development"
	LISTEN_ADDRESS = "0.0.0.0:3000"
)

var flagDebug = flag.Bool("debug", false, "Enable debug-level log")

func main() {
	flag.Parse()
	if *flagDebug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	config.Init()

	model.Init()
	controller.Init()

	err := controller.Engine.Run(LISTEN_ADDRESS)
	if err != nil {
		panic(xerrors.Errorf("error when opening controller: %w", err))
	}

	fmt.Printf("Server listening at %s", LISTEN_ADDRESS)
}
