package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Mostrar toda la información
func allDataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := readJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// añadir reporte
func addReporte(w http.ResponseWriter, r *http.Request) {
	var newReporte Reportes
	json.NewDecoder(r.Body).Decode(&newReporte)
	newReporte.ID = len(reporte) + 1
	reporte = append(reporte, newReporte)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newReporte)
}

// filtro por nivel de riesgo
func nivelHandler(w http.ResponseWriter, r *http.Request) {
	nivel := r.URL.Query().Get("nivel")
	data, err := readJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var filteredData []Reportes
	for _, d := range data {
		if d.Nivel == nivel {
			filteredData = append(filteredData, d)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredData)
}

// filtro por fecha del formulario
func formDateHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	data, err := readJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var filteredData []Reportes
	fromTime, _ := time.Parse("2006-01-02", from)
	toTime, _ := time.Parse("2006-01-02", to)
	for _, d := range data {
		if d.Fecha.After(fromTime) && d.Fecha.Before(toTime) {
			filteredData = append(filteredData, d)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredData)
}

// filtro por usuario un de estudiante
func estudianteHandler(w http.ResponseWriter, r *http.Request) {
	estudiante := r.URL.Query().Get("usuario_estudiante")
	data, err := readJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var filteredData []Reportes
	for _, d := range data {
		if d.Usuario_Estudiante == estudiante {
			filteredData = append(filteredData, d)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredData)
}

// filtro por usuario un de tutor
func tutorHandler(w http.ResponseWriter, r *http.Request) {
	tutor := r.URL.Query().Get("usuario_tutor")
	data, err := readJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var filteredData []Reportes
	for _, d := range data {
		if d.Usuario_Tutor == tutor {
			filteredData = append(filteredData, d)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredData)
}

// filtro por documento
func documentoHandler(w http.ResponseWriter, r *http.Request) {

	documento := r.URL.Query().Get("documento")
	doc, err := strconv.Atoi(documento)
	if err != nil {
		http.Error(w, "Documento invalido", http.StatusBadRequest)
		return
	}
	data, err := readJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var filteredData []Reportes
	for _, d := range data {
		if d.Documento == doc {
			filteredData = append(filteredData, d)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredData)
}
