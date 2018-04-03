package clcitybusapi

import "bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/geo"

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

// Recorrido represents the 'Recorrido' for a 'Linea' as of 'Cuando Llega City Bus' API.
type Recorrido struct {
	linea  *Linea
	puntos []*geo.Point
}

// Linea returns 'Linea' associated with this 'Recorrido'.
func (r *Recorrido) Linea() *Linea { return r.linea }

// Puntos returns all geo points of given 'Recorrido'.
func (r *Recorrido) Puntos() []*geo.Point { return r.puntos }

// NewRecorrido creates a new recorrido based on data passed.
func NewRecorrido(l *Linea, p []*geo.Point) *Recorrido {
	return &Recorrido{
		linea:  l,
		puntos: p,
	}
}

// Client is the interface that the client module should implement.
type Client interface {
	ParadaService() ParadaService
	LineaService() LineaService
	RecorridoService() RecorridoService
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

// RecorridoService has actions to fetch 'Recorrido' data from Cuando Llega City Bus API.
type RecorridoService interface {
	RecorridoDeLinea(l *Linea) (*Recorrido, error)
	RecorridosPorEmpresa(CodigoEmpresa int) ([]*Recorrido, error)
}
