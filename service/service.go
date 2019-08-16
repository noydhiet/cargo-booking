package service

import (
	"database/sql"
	// "fmt"
	"kit/log"
	"os"
	// "strings"
	_ "github.com/go-sql-driver/mysql" 

	dt "cargo-booking/datastruct"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "PASSWORD"
	dbName := "db_go"
	dbIP := "192.168.20.9"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

// func HelloWorld(name string) string {

// 	var helloOutput string

// 	helloOutput = fmt.Sprintf("Hello, $s ", name)

// 	return helloOutput

// }

// func HelloDaerah(name string, kelamin string, asal string) string {

// 	var helloOutput string

// 	switch strings.ToUpper(asal) {
// 	case "JAKARTA":
// 		{
// 			if strings.ToUpper(kelamin) == "PRIA" {
// 				helloOutput = fmt.Sprintf("Hi, MR ", name)
// 			} else {
// 				helloOutput = fmt.Sprintf("Hi, MRS  ", name)
// 			}
// 		}
// 	case "BANDUNG":
// 		{
// 			if strings.ToUpper(kelamin) == "PRIA" {
// 				helloOutput = fmt.Sprintf("Wilujeng, MR ", name)
// 			} else {
// 				helloOutput = fmt.Sprintf("Wilujeng, MRS  ", name)
// 			}
// 		}
// 	case "MEDAN":
// 		{
// 			if strings.ToUpper(kelamin) == "PRIA" {
// 				helloOutput = fmt.Sprintf("Horas, MR ", name)
// 			} else {
// 				helloOutput = fmt.Sprintf("Horas, MRS  ", name)
// 			}
// 		}

// 	}

	//helloOutput = fmt.Sprintf("Hello, $s ", name)

	// return helloOutput

// }

func InsertDelivery(del dt.Delivery) []dt.Delivery {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log("Checking Database")
	db := dbConn()
	// selDB, err := db.Query("INSERT * FROM t_mtr_delivery where route_id = ?", del.ID_ROUTE)
	selDB, err := db.Query("INSERT INTO t_trx_delivery(idt_delivery, routing_status,transport_status, last_known_location, fk_id_itenary, fk_id_route_spec) VALUES (?, ?, ?, ?, ?, ?)", del.ID_DELIVERY, del.ROUTING_STATUS, del.TRANSPORT_STATUS,del.LAST_KNOWN_LOCATION,del.ID_ITENARY,del.ID_ROUTE)

	if err != nil {
		panic(err.Error())
	}
	delv := dt.Delivery{}
	res := []dt.Delivery{}

	for selDB.Next() {
		var idDelivery, routeId, itenaryId int
		var routingStatus, transportStatus, lastKnownLoc string
		err = selDB.Scan(&idDelivery, &routingStatus, &transportStatus, &lastKnownLoc, &itenaryId, &routeId)
		if err != nil {
			panic(err.Error())
		}
		delv.ID_DELIVERY = idDelivery
		delv.ROUTING_STATUS = routingStatus
		delv.TRANSPORT_STATUS = transportStatus
		delv.LAST_KNOWN_LOCATION = lastKnownLoc
		delv.ID_ITENARY = itenaryId
		delv.ID_ROUTE = routeId
		res = append(res, delv)
	}

	return res

}
