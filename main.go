package main

import (
	_ "expvar"
	"net/http"
	"os"

	"cargo-booking/transport"

	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
)

// func dbcon() (db *sql.DB) {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := "PASSWORD"
// 	dbName := "db_go"
// 	dbIP := "192.168.20.9"
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

func main() {

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpsServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8778"
	}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}
