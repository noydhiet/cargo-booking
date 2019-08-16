package service

import (
	"database/sql"

	"os"

	// _ "github.com/go-sql-driver/mysql"
	dt "cargo-booking/datastruct"

	"github.com/go-kit/kit/log"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "PASSWORD"
	dbName := "db_go"
	dbIp := "192.168.20.9"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIp+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetRoute(del dt.Route) []dt.Route {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log("Checking Database")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM t_mtr_route_spec where route_id = ?", del.Id_route_spec)
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
		delv.Id_route_spec = id_route_spec
		delv.Duration = duration
		delv.Origin = origin
		delv.Destination = destination
		res = append(res, delv)
	}

	return res

}
