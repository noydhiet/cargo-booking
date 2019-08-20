package service

import (
	"database/sql"

	"os"

	dt "cargo-booking/datastruct"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kit/kit/log"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_go"
	dbIp := "127.0.0.1:3306"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIp+")/"+dbName)
	// db, err := sql.Open("mysql", "root: @tcp(127.0.0.1)/db_go")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetRoute(del dt.Route) []dt.Route {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log("Checking Database")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM t_mtr_spec where id_route_spec = ?", del.ID_ROUTE_SPEC)
	if err != nil {
		panic(err.Error())
	}
	delv := dt.Route{}
	res := []dt.Route{}

	for selDB.Next() {
		var id_route_spec, duration int
		var origin, destination string
		err = selDB.Scan(&id_route_spec, &duration, &origin, &destination)
		if err != nil {
			panic(err.Error())
		}
		delv.ID_ROUTE_SPEC = id_route_spec
		delv.DURATION = duration
		delv.ORIGIN = origin
		delv.DESTINATION = destination
		res = append(res, delv)
	}

	return res

}
