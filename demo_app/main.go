package main

import (
	_ "embed"
	"log"
	"os"
	"strings"

	"simple-go-auth/controller"
	"simple-go-auth/services/mysql"
)

//go:embed init/init_mysql.sql
var query string

func main() {
	if err := InitMysql(); err != nil && !strings.Contains(err.Error(), "Trigger already exists") {
		log.Fatalf("Can't init mysql %v", err)
	}
	Run()
}

func InitMysql() error {
	_, _, err := mysql.GetClientClient().ExecuteQuery(query)
	return err
}

func Run() {
	server := controller.Server{}
	appAddr := ":" + os.Getenv("PORT")

	server.Initialize()

	server.Run(appAddr)
}
