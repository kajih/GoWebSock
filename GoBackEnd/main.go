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
	log.Printf("WS Page Accessed: %+v\n\n", r)

	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	})

	if err != nil {
		log.Printf("websock error %v\n", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "websocket closed")

	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
	defer cancel()

	go func() {

		buf := make([]byte, 255)

		for ctx.Err() == nil {
			_, r, _ := c.Reader(ctx)
			n, readErr := r.Read(buf)
			if err != nil {
				log.Println(readErr.Error())
			}
			fmt.Printf("Recieved: [%s]\n", string(buf[:n]))
		}
	}()

	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			continue
		case t := <-ticker.C:
			err = wsjson.Write(ctx, c, t)
			if err != nil {
				log.Printf("Error Websocket Write %v\n", err.Error())
				cancel()
			}
			continue
		}
	}

	log.Println("Loop ended")
	log.Printf("Context was %v\n", ctx.Err().Error())
	c.Close(websocket.StatusNormalClosure, "")
}

// root http endpoint
func rootPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Root Page Accessed: %+v", r)
	enableCors(&w)
	fmt.Fprintf(w, "RESPONSE")
}
