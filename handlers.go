package main

import (
	"encoding/json"
	"net/http"
)

//Mostrar toda la información
func allDataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := readJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//filtro por nivel de riesgo
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
