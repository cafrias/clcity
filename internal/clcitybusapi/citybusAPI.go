package clcitybusapi

// Parada represents a 'Parada' as used by app 'Cuando Llega City Bus'.
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

// Linea represents a 'Linea' as used by app 'Cuando Llega City Bus'.
type Linea struct {
	CodigoLineaParada string
	Descripcion       string
	CodigoEntidad     string
	CodigoEmpresa     int
}

// Client is the interface that the client module should implement.
type Client interface {
	ParadaService() ParadaService
	LineaService() LineaService
}

// ParadaService represents a service for 'Parada'
type ParadaService interface {
	ParadasPorLinea(CodigoLineaParada string) ([]*Parada, error)
}

// LineaService has actions to fetch 'Linea' data from Cuando Llega City Bus API.
type LineaService interface {
	LineasPorEmpresa(CodigoEmpresa int) ([]*Linea, error)
}
