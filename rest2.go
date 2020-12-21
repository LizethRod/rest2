package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

type Contacto struct {
	ID_Contacto    int    `json:"id"`
	NombreCompleto string `json:"nombre"`
	Email          string
	TelMovil       string
	FechaCaptura   string
	TemaDeInteres  *TemaDeInteres
	Mensaje        string
	Estado         *Estado
}

type TemaDeInteres struct {
	ID_TemaDeInteres int
	Nombre           string
}

type Estado struct {
	ID_Estado string
	Nombre    string
}

var (
	contador  = 0
	contactos []Contacto
)

func main() {
	fmt.Println("holaa 7u7")
	//http.HandleFunc("/", LandingpPage)
	contactos = append(contactos, Contacto{ID_Contacto: 1, NombreCompleto: "Lizeth", TemaDeInteres: &TemaDeInteres{ID_TemaDeInteres: 1, Nombre: "Contaminación"}})
	contactos = append(contactos, Contacto{ID_Contacto: 2, NombreCompleto: "Juan", TemaDeInteres: &TemaDeInteres{ID_TemaDeInteres: 2, Nombre: "Cursos"}})
	contactos = append(contactos, Contacto{ID_Contacto: 3, NombreCompleto: "Maria", Estado: &Estado{ID_Estado: "1", Nombre: "Aguascalientes"}})

	router := mux.NewRouter()
	router.HandleFunc("/", LandingPage).Methods("GET")
	router.HandleFunc("/contactos", GetContactosEndPoint).Methods("GET")
	router.HandleFunc("/contactosRegistrados", registrosExistentes).Methods("GET")
	router.HandleFunc("/contactos", GetContactosEndPoint).Methods("POST")
	router.HandleFunc("/contacto", PostContactoEndPoint)

	handlerCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)

	http.ListenAndServe("localhost:3030", handlerCORS)
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request")
	fmt.Fprintf(w, "<html><body>Petición recibida</body></html>")
	contador++
	fmt.Println(contador)
}

func GetContactosEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reporte de contactos")
	json.NewEncoder(w).Encode(contactos)
}

func GetContactoEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	indice, _ := strconv.Atoi(params["idx"])

	json.NewEncoder(w).Encode(contactos[indice])
}

func PostContactoEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Mensaje POST recibido")

	var nuevo Contacto
	json.NewDecoder(r.Body).Decode(&nuevo)
	contactos = append(contactos, nuevo)

	json.NewEncoder(w).Encode(contactos)
}

func registrosExistentes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Contactos registrados: ", len(contactos))
}
