package main

import (
	"github.com/immanoj16/taskbook/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	server, err := app.NewServer()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	err = server.Start()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	defer func() {
		if err = server.Shutdown(); err != nil {
			logrus.Fatal(err.Error())
		}
	}()
}
