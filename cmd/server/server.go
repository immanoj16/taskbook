package main

import (
	"github.com/immanoj16/taskbook/internal/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	s, err := server.NewServer()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
	err = s.DB.Start()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
	err = s.Start()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
	defer func() {
		if err = s.Shutdown(); err != nil {
			logrus.Fatal(err.Error())
		}
	}()
}
