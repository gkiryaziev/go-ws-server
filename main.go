package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	hdlLog "./handlers/log_handler"
	hdlTopic "./handlers/topic_handler"
	hdlWs "./handlers/ws_handler"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mx := mux.NewRouter()

	// variables
	hostIP := "192.168.2.22"
	hostPort := "8080"

	// http server address and port
	hostBind := fmt.Sprintf("%s:%s", hostIP, hostPort)

	// open connection to database
	db, err := sqlx.Connect("sqlite3", "base.db3")
	checkError(err)
	db.SetMaxIdleConns(100)
	defer db.Close()

	// websocket hub
	hub := hdlWs.NewHub(db)
	go hub.Run()

	// controllers
	wsCtrl := hdlWs.NewWsController(hub)
	logCtrl := hdlLog.NewLogController(db)
	topicCtrl := hdlTopic.NewTopicController(db)

	// user handler
	mx.HandleFunc("/ws", wsCtrl.WsHandler)
	mx.HandleFunc("/api/v1/logs", logCtrl.GetLogs).Methods("GET")
	mx.HandleFunc("/api/v1/topics", topicCtrl.GetTopics).Methods("GET")

	// static
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	// start http server
	fmt.Println("Listening on", hostBind)
	err = http.ListenAndServe(hostBind, mx)
	checkError(err)
}
