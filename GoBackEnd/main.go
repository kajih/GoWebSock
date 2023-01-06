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

// CORS control for http endpoints (not websock)
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Websocket endpoint
func ws(w http.ResponseWriter, r *http.Request) {
	log.Printf("WS Page Accessed: %+v", r)

	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	})

	if err != nil {
		log.Printf("websock error %v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "websocket closed")

	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
	defer cancel()

	loop := true
	for loop {
		select {
		case <-ctx.Done():
			loop = false
			continue
		case t := <-ticker.C:
			err = wsjson.Write(ctx, c, t)
			if err != nil {
				log.Printf("Error Websocket Write %v", err.Error())
				loop = false
			}
			continue
		}
	}

	log.Println("Loop ended")
	c.Close(websocket.StatusNormalClosure, "")
}

// root http endpoint
func rootPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Root Page Accessed: %+v", r)
	enableCors(&w)
	fmt.Fprintf(w, "RESPONSE")
}
