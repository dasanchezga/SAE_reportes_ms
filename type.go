package main

type Reportes struct{
	//Formularios
	ID int `json:"id"`
	Nombre_Formulario string `json:"nombre_form"`
	Nivel string `json:"nivel"`
	Documento int64 `json:"documento"`
	Tipologia string `json:"tipologia"`
	
	//Tutorias
	Usuario_Estudiante string `json:"usuario_un_estudiante"`
	Usuario_Tutor string `json:"usuario_un_tutor"`
	Fecha string `json:"fecha"`
	Objetivo string `json:"objetivo"`
	Acuerdo string `json:"acuerdo"`

	//Observaciones 
	Fecha_Obs string `json:"fecha_obs"`
	Descripcion string `json:"descripcion"`

	//Remisiones
	Usuario_Estudiante2 string `json:"id_estudiante"` 
	Usuario_Tutor2 string `json:"id_docente"`
	Tipo_Remision string `json:"tipo_remision"`
	Estado_Remision string `json:"estado_remision"`
	Estado_Solicitud string `json:"estado"`
	Estado_PE string `json:"estado_pe"`
	Fecha_PE string `json:"fecha_primera_escucha"`
	Observacion_PE string `json:"observacion"`
	Justificacion string `json:"justificacion"`
	Fecha_Solicitud string `json:"fecha_solicitud_remision"`
	Fecha_Aprobacion string `json:"fecha_envio_remision"`
}

var reporte  []Reportes