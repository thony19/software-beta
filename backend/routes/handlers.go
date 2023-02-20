package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello API CHI"))
	if err != nil {
		log.Panicln(err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Tipo de dato
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newTask user
	resqBody, err := ioutil.ReadAll(r.Body) //informacion  del cliente

	if err != nil {
		fmt.Fprint(w, "Insert a Valid user")
	}

	// Pasando los datos a la variable
	json.Unmarshal(resqBody, &newTask)

	//generando ID para la nueva tarea
	newTask.ID = len(users) + 1

	// Guardando en la struct
	users = append(users, newTask)

	// Respondo a la tarea
	w.Header().Set("Content-Type", "application/json") //Tipo de dato
	w.WriteHeader(http.StatusCreated)                  //todo ha ido bien y el dato asignado a la lista de datos
	json.NewEncoder(w).Encode(newTask)                 //Enviando dato
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")     //Extrae las variables del request
	taskID, err := strconv.Atoi(id) //variable de la ruta. Convirtiendo a entero.

	if err != nil {
		fmt.Fprint(w, "Invalid ID")
		return
	}

	// Buscaremos si el ID que ingresaron está en la lista de tarea
	for _, user := range users { // Recorriendo la lista de tareas
		if user.ID == taskID {
			w.Header().Set("Content-Type", "application/json") //Tipo de dato
			json.NewEncoder(w).Encode(user)                    // Enviando la tarea del id
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")     //Extrae las variables del request
	userID, err := strconv.Atoi(id) //variable de la ruta. Convirtiendo a entero.

	if err != nil {
		fmt.Fprint(w, "Invalid ID")
		return
	}

	// Buscaremos si el ID que ingresaron está en la lista de tarea
	for index, user := range users { // Recorriendo la lista de tareas
		if user.ID == userID {
			users = append(users[:index], users[index+1:]...)
			fmt.Fprintf(w, "The task with ID %v has beenremove succesfully", userID) // Enviando la tarea del id
		}
	}
}
