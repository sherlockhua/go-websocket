# Go-WebSocket 

This package was replaced by the [Gorilla
websocket](https://github.com/gorilla/websocket) package. Users of this
package are encouraged to update to the Gorilla package. 

The code for this package is still available in the [master
branch](https://github.com/garyburd/go-websocket/tree/master/) for those who do
not want to update.

There are a few differences between the Gorilla package and this package:

- Pongs are handled using a callback.
- The request header argument to the Upgrade function is replaced with an http.Request argument.
- OpXXX constants renamed to XXXMessage.
