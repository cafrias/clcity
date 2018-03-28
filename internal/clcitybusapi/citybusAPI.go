package clcitybusapi

// Parada represents a 'Parada' as used by app 'Cuando Llega City Bus'.
type Parada struct {
	Codigo                     string
	Identificador              string
	Descripcion                string
	AbreviaturaBandera         string
	AbreviaturaAmpliadaBandera string
	LatitudParada              string
	LongitudParada             string
	AbreviaturaBanderaGIT      string
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
	ParadasPorLinea(CodigoLineaParada int) ([]*Parada, error)
	ParadasPorEmpresa(CodigoEmpresa int) ([]*Parada, error)
}

// LineaService has actions to fetch 'Linea' data from Cuando Llega City Bus API.
type LineaService interface {
	LineasPorEmpresa(CodigoEmpresa int) ([]*Linea, error)
}
