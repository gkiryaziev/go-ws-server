##	Golang Publish/Subscribe Websocket Server

[Go](https://golang.org/) websocket server example with [Gorilla](http://www.gorillatoolkit.org/) toolkit and [SQLite](https://www.sqlite.org/) database.
In this example you can subscribe, unsubscribe and publish messages.

ACTION - `SUBSCRIBE`, `UNSUBSCRIBE`, `PUBLISH`

Message example:
```
{"action" : "ACTION", "topic" : "TOPIC NAME", "data" : "DATA"}
```

PS Version with goroutines and channels.

![Mind](/mind.png?raw=true "Mind")