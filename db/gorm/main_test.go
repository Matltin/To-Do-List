package db

import (
	"log"
	"os"
	"testing"
	"to_do_list/db"
	"to_do_list/util"
)

var p db.DataBase

func TestMain(m *testing.M) {

	util.Load("../.env")

	p = NewPostgres()

	dsn := os.Getenv("DSN")

	log.Println("dsn ----> ", dsn)

	err := p.Connect(dsn)
	if err != nil {
		log.Fatal("Failed to connect db: ", err)
	}

	err = p.Init()
	if err != nil {
		log.Fatal("Fialed to init db: ", err)
	}

	os.Exit(m.Run())
}
