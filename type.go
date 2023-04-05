package main

import "time"

type Lista struct {
	Lugar    string `json:"lugar"`
	Estado   string `json:"estado"`
	Objetivo string `json:"objetivo"`
	Acuerdo  string `json:"acuerdo"`
}
type Observaciones struct {
	Fecha_Obs   string `json:"fecha_obs"`
	Descripcion string `json:"descripcion"`
}

type Remisiones struct{
	Usuario_Estudiante1 string `json:"usuario_estudiante1"`
	Usuario_Tutor1      string `json:"usuario_tutor1"`
	Tipo_Remision       string `json:"tipo_remision"`
	Estado_Remision     string `json:"estado_remision"`
	Estado_Solicitud    string `json:"estado_solicitud"`
	Estado_PE           string `json:"estado_PE"`
	Fecha_PE            string `json:"fecha_PE"`
	Observacion_PE      string `json:"observacion_PE"`
	Justificacion       string `json:"justificacion"`
	Fecha_Solicitud     string `json:"fecha_solicitud"`
	Fecha_Aprobacion    string `json:"fecha_aprobacion"`

}

type Reportes struct {
	//Usuarios
	Usuario_Estudiante0 string `json:"usuario_estudiante0"`
	Usuario_Tutor0      string `json:"usuario_tutor0"`
	Programa            string `json:"programa"`
	Departamento        string `json:"departamento"`
	Facultad            string `json:"facultad"`

	//Formularios
	ID                int    `json:"id"`
	Nombre_Formulario string `json:"nombre_form"`
	Nivel             string `json:"nivel"`
	Documento         int    `json:"documento"`
	Tipologia         string `json:"tipologia"`

	//Tutorias
	Usuario_Estudiante string    `json:"usuario_estudiante"`
	Usuario_Tutor      string    `json:"usuario_tutor"`
	Estado_Tutor       string    `json:"estado_tutor"`
	Fecha              time.Time `json:"fecha"`

	Lista []Lista `json:"lista"`

	//Observaciones
	Observaciones []Observaciones `json:"observaciones"`

	//Remisiones
	Remisiones []Remisiones `json:"remisiones"`
}

var reporte []Reportes
