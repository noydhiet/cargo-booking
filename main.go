package main

import (
	"cargo-booking/transport"
	_ "expvar"

	// "log"

	// "golang-service2/transport"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

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

	// log.Println("Server started on : http//localhost:8778")
}
