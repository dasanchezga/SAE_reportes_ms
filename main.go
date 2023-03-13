package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//router para los endpoints
	r := mux.NewRouter()
	r.HandleFunc("/repo", getReporte).Methods("GET")
	r.HandleFunc("/repo", addReporte).Methods("POST")
	r.HandleFunc("/repo/data", allDataHandler)
	r.HandleFunc("/repo/data/nivel", nivelHandler)
	r.HandleFunc("/repo/data/fecha", formDateHandler)
	r.HandleFunc("/repo/data/estudiante", estudianteHandler)
	r.HandleFunc("/repo/data/tutor", tutorHandler)
	r.HandleFunc("/repo/data/documento", documentoHandler)
	//r.HandleFunc("/report", updateReporte)
	//r.HandleFunc("/reporte", deleteReporte)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
