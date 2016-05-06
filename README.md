##	Golang Publish/Subscribe Websocket Server

[Go](https://golang.org/) websocket server with [Gorilla](http://www.gorillatoolkit.org/) toolkit and [SQLite](https://www.sqlite.org/) database.
With this server you can subscribe, unsubscribe and publish messages.

ACTION - `SUBSCRIBE`, `UNSUBSCRIBE`, `PUBLISH`

### Message example:
```
{"action" : "ACTION", "topic" : "TOPIC NAME", "data" : "DATA"}
```

![Mind](/mind.png?raw=true "Mind")

### Installation:
```
go get github.com/gkiryaziev/go-ws-server
```

### Edit configuration:
```
Copy `config.default.yaml` to `config.yaml` and edit configuration.
```

### Build and Run:
```
go build && go-ws-server
```
