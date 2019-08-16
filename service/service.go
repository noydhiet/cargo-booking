package service

import (
	dt "cargo-booking/datastruct"
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
func GetItenary(tam dt.Tampil) []dt.Tampil {
	db := dbcon()
	selDb, err := db.Query("SELECT * FROM t_mtr_itenary")

	if err != nil {
		panic(err.Error())
	}

	Ez := dt.Tampil{} //single object employee struct.Created to capture row

	res := []dt.Tampil{} //array of object Employee.create capture whole

	for selDb.Next() {
		var iditenary, voyagenumber, loadtime, unloadtime int
		var unloadlocation, loadlocation string
		err = selDb.Scan(&iditenary, &loadlocation, &unloadlocation, &loadtime, &unloadtime, &voyagenumber)
		if err != nil {
			panic(err.Error())
		}
		Ez.ID_ITENARY = iditenary
		Ez.LOAD_LOCATION = loadlocation
		Ez.UNLOAD_LOCATION = unloadlocation
		Ez.UNLOAD_TIME = unloadtime
		Ez.LOAD_TIME = loadtime
		Ez.VOYAGE_NUMBER = voyagenumber
		res = append(res, Ez)
	}
	return res
}
