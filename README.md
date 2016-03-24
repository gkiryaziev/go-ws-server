##	Golang Publish/Subscribe Websocket Server

[Go](https://golang.org/) websocket server with [Gorilla](http://www.gorillatoolkit.org/) toolkit and [SQLite](https://www.sqlite.org/) database.
With this server you can subscribe, unsubscribe and publish messages.

ACTION - `SUBSCRIBE`, `UNSUBSCRIBE`, `PUBLISH`

Message example:
```
{"action" : "ACTION", "topic" : "TOPIC NAME", "data" : "DATA"}
```

![Mind](/mind.png?raw=true "Mind")

## Installation

#### 1. Install GO
#### 2. Install GB
  `go get -u github.com/constabulary/gb/...`
#### 3. Clone project
  `git clone https://github.com/gkiryaziev/go_gorilla_pubsub_websocket_server.git`
#### 4. Restore vendors
  `cd go_gorilla_pubsub_websocket_server`
  
  `gb vendor restore`
#### 5. Edit configuration
  Copy `config.default.yaml` to `config.yaml` and edit configuration.
#### 6. Build and Run project
  `gb build && bin/ws.server run`