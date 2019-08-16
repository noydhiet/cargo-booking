package service

import (
	"cargo-booking/datastruct"
	"database/sql"
	"fmt"
)

func dbcon() (db *sql.DB) {
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

func HelloWorld(name string) string {
	var helloOutput string
	helloOutput = fmt.Sprintf("Hello, %s ", name)
	return helloOutput
}

func HelloDaerah(name, jenis_kelamin, asal_kota string) string {
	var helloOutput string
	if (jenis_kelamin == "wanita") || (jenis_kelamin == "Wanita") || (jenis_kelamin == "WANITA") || (jenis_kelamin == "w") || (jenis_kelamin == "W") {
		jenis_kelamin = "MRS."
	} else if (jenis_kelamin == "pria") || (jenis_kelamin == "Pria") || (jenis_kelamin == "PRIA") || (jenis_kelamin == "p") || (jenis_kelamin == "P") {
		jenis_kelamin = "MR."
	} else {
		jenis_kelamin = "MR/MRS."
	}

	if (asal_kota == "jakarta") || (asal_kota == "Jakarta") || (asal_kota == "JAKARTA") {
		asal_kota = "Hi"
	} else if (asal_kota == "bandung") || (asal_kota == "Bandung") || (asal_kota == "BANDUNG") {
		asal_kota = "Wilujeung"
	} else if (asal_kota == "medan") || (asal_kota == "Medan") || (asal_kota == "MEDAN") {
		asal_kota = "Horas"
	} else {
		asal_kota = "Gagal"
	}

	helloOutput = fmt.Sprintf("%s %s %s", asal_kota, jenis_kelamin, name)
	return helloOutput
}
func GetItenary(voyage_number int, unload_location, load_location string) string {
	var GetItenary string
	db := dbcon()
	selDb, err := db.Query("SELECT * FROM t_mtr_itenary ORDER BY id_itenary ASC")

	if err != nil {
		panic(err.Error())
	}

	Ez := datastruct.GetItenaryRequest{} //single object employee struct.Created to capture row

	//res := datastruct.GetItenaryResponse{} //array of object Employee.create capture whole

	for selDb.Next() {
		var id_itenary, voyage_number, load_time, unload_time int
		var unload_location, load_location string
		err = selDb.Scan(&id_itenary, &voyage_number, &unload_location, &load_location, &load_time, &unload_time)
		if err != nil {
			panic(err.Error())
		}
		Ez.VOYAGE_NUMBER = voyage_number
		Ez.LOAD_LOCATION = load_location
		Ez.UNLOAD_LOCATION = unload_location
		//println("ID", id, "NAME", name, "CITY", city)
		//res = append(res, emp)
	}
	defer db.Close()
	GetItenary = fmt.Sprintf("%T %T %T ", voyage_number, unload_location, load_location)
	return GetItenary
}
