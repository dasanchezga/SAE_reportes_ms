package main

import (
    "encoding/json"
    "net/http"
)

func getReporte(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(reporte)
}

func addReporte(w http.ResponseWriter, r *http.Request) {
	var newReporte Reportes
	json.NewDecoder(r.Body).Decode(&newReporte)
	newReporte.ID = len(reporte) + 1
	reporte = append(reporte, newReporte)
	json.NewEncoder(w).Encode(newReporte)
}

/*func updateReporte(w http.ResponseWriter, r *http.Request) {
	var updatedReporte Reportes
	json.NewDecoder(r.Body).Decode(&updatedReporte)
	for i, report := range reporte {
		if report.ID == updatedReporte.ID {
			reporte[i] = updatedReporte
			break
		}
	}
	json.NewEncoder(w).Encode(updatedReporte)
}

func deleteReporte(w http.ResponseWriter, r *http.Request) {
	var deletedReporte Reportes
	json.NewDecoder(r.Body).Decode(&deletedReporte)
	for i, report:= range reporte {
		if report.ID == deletedReporte.ID {
			reporte = append(reporte[:i], reporte[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(deletedReporte)
}*/