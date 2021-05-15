package main

import (
	"github.com/immanoj16/taskbook/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	err = s.Start()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	defer func() {
		if err = s.Shutdown(); err != nil {
			logrus.Fatal(err.Error())
		}
	}()
}
