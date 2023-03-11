package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//lee el JSON que recibe la API
// "./ext/generated.json"
func readJSON() ([]Reportes, error) {
	file, err := os.Open("./ext/generated.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return reporte, err
	}
	defer file.Close()
	decodificador := json.NewDecoder(file)
	err = decodificador.Decode(&reporte)
	if err != nil {
		fmt.Println("Error al analizar el archivo JSON:", err)
		return reporte, err
	}
	return reporte, nil
}
