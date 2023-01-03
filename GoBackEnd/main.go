package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	fmt.Println("Server Started ***")
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/ws", ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func ws(w http.ResponseWriter, r *http.Request) {
	log.Printf("WS Page Accessed: %+v", r)
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("websock error %v", err)
	}
	defer c.Close(websocket.StatusInternalError, "websocket closed")

	tick := time.Tick(500 * time.Millisecond)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	for v := range tick {
		err = wsjson.Write(ctx, c, v)
		if err != nil {
			log.Printf("Error Websocket Write %v", v)
		}
	}

	var v interface{}
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		log.Printf("websock error %v", err)
	}

	log.Printf("received: %v", v)

	c.Close(websocket.StatusNormalClosure, "")
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Root Page Accessed: %+v", r)
	enableCors(&w)
	fmt.Fprintf(w, "RESPONSE")
}
