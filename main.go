package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sahilgarg96/DBTNT/handler"
	"github.com/sahilgarg96/DBTNT/logging"
	"github.com/sahilgarg96/DBTNT/redis"
	"github.com/sahilgarg96/DBTNT/scheduler"
	"log"
	"net/http"
)

var Logger = logging.NewLogger()

func main() {

	Logger.Infof("starting service ")

	redis.Init()
	scheduler.Init()

	router := mux.NewRouter()
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods(http.MethodPost).Path("/sendPdf/{user_id:[0-9]+}").HandlerFunc(handler.GeneratePdf)

	err := http.ListenAndServe(":8080", router)
	fmt.Println(err)
	if err != nil {
		log.Fatal("some error occured")
	}

}
