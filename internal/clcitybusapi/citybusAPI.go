package clcitybusapi

// Parada representa una parada como es usado por Cuando Llega.
type Parada struct {
	Codigo                     string `json:"Codigo"`
	Identificador              string `json:"Identificador"`
	Descripcion                string `json:"Descripcion"`
	AbreviaturaBandera         string `json:"AbreviaturaBandera"`
	AbreviaturaAmpliadaBandera string `json:"AbreviaturaAmpliadaBandera"`
	LatitudParada              string `json:"LatitudParada"`
	LongitudParada             string `json:"LongitudParada"`
	LongitudParadaGIT          string `json:"LongitudParadaGIT"`
}

// Linea representa una linea como es usada por Cuando Llega.
type Linea struct {
	CodigoLineaParada string `json:"CodigoLineaParada"`
	Descripcion       string `json:"Descripcion"`
	CodigoEntidad     string `json:"CodigoEntidad"`
	CodigoEmpresa     string `json:"CodigoEmpresa"`
}
