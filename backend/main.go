package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	socketio "github.com/googollee/go-socket.io"
)

type user struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allUsers []user

var users = allUsers{
	{
		ID:      1,
		Name:    "Thony Javier",
		Content: "Programador fullstack",
	},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Tipo de dato
	json.NewEncoder(w).Encode(users)
}

func main() {
	// INCIAR EL SERVIDOR
	svr := socketio.NewServer(nil)

	svr.OnConnect("/", func(so socketio.Conn) error { // este es el cliente conetado
		so.SetContext("")
		so.Join("chat_room") // Insertando cada cliente en uns sala aparte
		fmt.Println("A new user connected: ", so.ID())

		return nil
	})

	// svr.BroadcastToRoom("chat_room", "chat message", msg)
	svr.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		svr.BroadcastToRoom("", "chat_room", "chat message", msg) // Retrasmitiendo a todos los que esten en la sala
		fmt.Println(msg)
	})

	svr.OnDisconnect("/", func(so socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go svr.Serve()
	defer svr.Close()

	// GO-CHI
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Creando Rutas
	r.Handle("/socket.io/", svr)
	r.Handle("/", http.FileServer(http.Dir("../frontend/public")))
	r.Get("/users", GetUsers)

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Panic(err)
	}

}
