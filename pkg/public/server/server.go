package server

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"github.com/gorilla/mux"
	"github.com/alexbt/go-mock-server/pkg/internal/api"
)

var contextRoot = "/{any:.*}"



func StartServer() {
	port := getPort()
	router := mux.NewRouter()
	router.HandleFunc(contextRoot, api.HandleRequests)

	httpServer := http.ListenAndServe(fmt.Sprintf(":%s", port), router)

	log.Fatal(httpServer)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println(port)
	if (len(os.Args) > 1) {
		port = os.Args[1]
	}
	return port
}
