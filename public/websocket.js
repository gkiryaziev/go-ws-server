var output;
var websocket;
var connected = false;

function init(){
    output = document.getElementById("output");
}

function connect() {
    if ("WebSocket" in window) {
        if (!connected) {
            var address = document.getElementById('txtAddress').value;
            websocket = new WebSocket(address);
            websocket.onopen = onOpen;
            websocket.onclose = onClose;
            websocket.onmessage = onMessage;
            websocket.onerror = onError;
        }
    } else {
        alert("WebSockets not supported on your browser.");
    }
}

function disconnect() {
    websocket.close()
}

function sendMessage() {
    if (connected) {
        var data = document.getElementById('txtText').value;
        websocket.send(data);
        output.innerHTML += "<<== " + data + "<br>";
    }
}

function onOpen(event) {
    connected = true;
}
function onClose(event) {
    if (event.wasClean) {
        output.innerHTML += "Соединение закрыто чисто<br>";
    } else {
        output.innerHTML += "Обрыв соединения<br>";
    }
    output.innerHTML += "Код: " + event.code + " причина: " + event.reason + "<br>";
    connected = false;
}

function onError(error) {
    output.innerHTML += "Ошибка " + error.message + "<br>";
    connected = false;
}

function onMessage(e) {
    output.innerHTML += "==>> " + e.data + "<br>";
}

// ---------- pub/sub

function subscribe() {
    sendData("SUBSCRIBE");
}

function unsubscribe() {
    sendData("UNSUBSCRIBE");
}

function publish() {
    sendData("PUBLISH");
}

function sendData(action) {
    if (connected) {
        var topic = document.getElementById('txtTopic').value;
        var data = document.getElementById('txtData').value;
        var msg = "{\"action\" : \""+ action +"\", \"topic\" : \""+ topic +"\", \"data\" : \""+ data +"\"}";
        websocket.send(msg);
        output.innerHTML += "<<== " + msg + "<br>";
    }
}