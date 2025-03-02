package main

import (
	// "database/sql"
	// "log"
	// "os"
	// "to_do_list/api"
	// "to_do_list/db/sqlc"
	// "to_do_list/token"
	// "to_do_list/util"

	_ "github.com/lib/pq"
)

func main() {
	// util.Load(".env")
	// db := db.NewPostgres()

	// dsn := os.Getenv("DSN")
	// err := db.Connect(dsn)
	// if err != nil {
	// 	log.Fatal("can not connect to db: ", err)
	// }

	// err = db.Init()
	// if err != nil {
	// 	log.Fatal("can not initiliaze the db: ", err)
	// }

	// connStr := "postgresql://root:secret@localhost:5432/tddb?sslmode=disable"

	// Open a connection to the database
	// conn, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	// store := sqlc.NewStore(conn)

	// symmetric := os.Getenv("PASETO_SYMMETRIC_KEY")

	// maker, err := token.NewPasetoMaker(symmetric)
	// if err != nil {
	// 	log.Fatal("can not make new Paseto: ", err)
	// }

	// serverAddress := os.Getenv("SERVER_ADDRESS")

	// server := api.NewServer(store, *maker)
	// server.Start(serverAddress)
}
