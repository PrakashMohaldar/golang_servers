package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn	

func main(){
	// create endpoint

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request){
		// conn is a pointer to websocket
		conn, _ := upgrader.Upgrade(w, r,nil)
		clients = append(clients, *conn)

		// loop if client send to server
		for{
			// read from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return 
			}

			// print message in console
			fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))

			// loop if message found and send again to client for write in browser
			for _,clientconn := range clients{
				if err := clientconn.WriteMessage(msgType, msg); err!=nil{
					return
				}
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"index.html")
	})
	println("server listening to port 8080")
	http.ListenAndServe(":8080", nil)
}
