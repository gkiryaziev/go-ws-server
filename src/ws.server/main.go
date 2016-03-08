package main

import (
	"fmt"
	"log"
	"net/http"

	hdlLog "ws.server/handlers/log_handler"
	hdlSubscriber "ws.server/handlers/subscriber_handler"
	hdlTopic "ws.server/handlers/topic_handler"
	hdlWs "ws.server/handlers/ws_handler"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mx := mux.NewRouter()

	// variables
	hostIP := "127.0.0.1"
	hostPort := "8000"

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
	subscriberCtrl := hdlSubscriber.NewSubscriberController(db)

	// user handler
	mx.HandleFunc("/ws", wsCtrl.WsHandler)
	mx.HandleFunc("/api/v1/logs", logCtrl.GetLogs).Methods("GET")
	mx.HandleFunc("/api/v1/topics", topicCtrl.GetTopics).Methods("GET")
	mx.HandleFunc("/api/v1/subscribers", subscriberCtrl.GetSubscribers).Methods("GET")

	// static
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	// start http server
	fmt.Println("Listening on", hostBind)
	err = http.ListenAndServe(hostBind, mx)
	checkError(err)
}