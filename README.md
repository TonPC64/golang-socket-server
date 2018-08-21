# GO SOCKET SERVER

```js
let ws = new WebSocket('ws://localhost:8000/chann')
ws.onopen = function (evt) {
  console.log('OPEN')
}
ws.onclose = function (evt) {
  console.log('CLOSE')
  ws = null
}
ws.onmessage = function (evt) {
  console.log('RESPONSE: ' + evt.data)
}
ws.onerror = function (evt) {
  console.log('ERROR: ' + evt.data)
}
```