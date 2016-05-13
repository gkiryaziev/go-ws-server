##	Golang Publish/Subscribe Websocket Server

[![Go Report Card](https://goreportcard.com/badge/github.com/gkiryaziev/go-ws-server)](https://goreportcard.com/report/github.com/gkiryaziev/go-ws-server)

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

### Packages:
You can use [glide](https://glide.sh/) packages manager to get all needed packages.
```
go get -u -v github.com/Masterminds/glide

cd go-ws-server && glide install
```