package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	psqlDBOnce sync.Once
	psqlConn   *sql.DB
)

func PsqlInit(dbConnectionString string) {
	psqlDBOnce.Do(func() {
		log.Println("Initializing psql database connection...")
		var err error

		psqlConn, err = sql.Open("postgres", dbConnectionString)
		if err != nil {
			log.Fatal("Unable to open connection to the postgres server: ", err)
		}

		if err := psqlConn.Ping(); err != nil {
			log.Fatal("Unable to reach the postgres server: ", err)
		}
	})

}

func PsqlClient() *sql.DB {
	return psqlConn
}
