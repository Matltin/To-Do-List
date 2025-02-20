package main

import (
	"log"
	"os"
	"to_do_list/api"
	"to_do_list/db"
	"to_do_list/token"
	"to_do_list/util"
)

func main() {
	util.Load(".env")
	db := db.NewPostgres()

	dsn := os.Getenv("DSN")
	err := db.Connect(dsn)
	if err != nil {
		log.Fatal("can not connect to db: ", err)
	}

	err = db.Init()
	if err != nil {
		log.Fatal("can not initiliaze the db: ", err)
	}

	symmetric := os.Getenv("PASETO_SYMMETRIC_KEY")

	maker, err := token.NewPasetoMaker(symmetric)
	if err != nil {
		log.Fatal("can not make new Paseto: ", err)
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")

	server := api.NewServer(db, *maker)
	server.Start(serverAddress)
}