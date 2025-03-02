package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"to_do_list/util"

	_ "github.com/lib/pq"
)

var testStore Store


func TestMain(m *testing.M) {
	util.Load("../../.env")

	dbSource := "postgresql://root:secret@localhost:5432/tddb?sslmode=disable"

	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("can not connect to db: ", err)
	}

	testStore = NewStore(conn)

	os.Exit(m.Run())
}
