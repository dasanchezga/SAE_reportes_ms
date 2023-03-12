package main

import (
	"encoding/json"
	"net/http"
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
