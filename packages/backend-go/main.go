// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"simplechatagent-go/controllers"
	"simplechatagent-go/initializers"
	"simplechatagent-go/ws"
)

var PORT = os.Getenv("PORT")

var addr = flag.String("addr", ":"+PORT, "http service address")

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	// initializers.SyncDB()
}

func main() {
	log.Println("go run on port:", PORT)
	flag.Parse()
	hub := ws.NewHub()
	go hub.Run()
	http.HandleFunc("/", controllers.ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})

	http.HandleFunc("/c", controllers.PostsIndex)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
