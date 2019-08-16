package main

import (
	_ "expvar"
	"cargo-booking/transport"
	"kit/log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql" 
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	transport.RegisterHttpsServicesAndStartListener()
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8080"
	//}
	port := "8778"
	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
		//fmt.Println("Error")
	}
}
