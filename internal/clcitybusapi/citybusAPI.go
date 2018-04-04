package clcitybusapi

import "bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/geo"

// Empresa represents a transport agency.
type Empresa struct {
	Codigo  int
	Nombre  string
	URL     string
	Lang    string
	TZ      string
	Lineas  []*Linea
	Paradas []*Parada
}

// NewEmpresa creates a new 'Empresa' with defaults.
func NewEmpresa(cod int) *Empresa {
	return &Empresa{
		Codigo: cod,
		Nombre: "City Bus",
		URL:    "https://riogrande.gob.ar/",
		Lang:   "SPA",
		TZ:     "America/Argentina/Ushuaia",
	}
}

// Parada represents a stop related with a 'Empresa'.
type Parada struct {
	Codigo string
	Punto  geo.Point
}

// ParadaLinea represents a stop for a 'Linea'.
type ParadaLinea struct {
	Codigo                     int
	Identificador              string
	Descripcion                string
	AbreviaturaBandera         string
	AbreviaturaAmpliadaBandera string
	AbreviaturaBanderaGIT      string
	Punto                      geo.Point
	Linea                      *Linea
}

// Linea represents a route for a given 'Empresa'.
type Linea struct {
	Codigo        int
	Descripcion   string
	CodigoEntidad int
	Empresa       *Empresa
	Paradas       []*ParadaLinea
	Recorrido     *Recorrido
}

// Recorrido represents the shape for 'Linea' to be draw on a plane.
type Recorrido struct {
	puntos []geo.Point
}

// Puntos returns all geo points of given 'Recorrido'.
func (r *Recorrido) Puntos() []geo.Point { return r.puntos }

// NewRecorrido creates a new recorrido based on data passed.
func NewRecorrido(l *Linea, p []geo.Point) *Recorrido {
	return &Recorrido{
		puntos: p,
	}
}

// Client is the interface that the client module should implement.
type Client interface {
	ParadaService() ParadaService
	LineaService() LineaService
	RecorridoService() RecorridoService
	EmpresaService() EmpresaService
}

// ParadaService represents a service for 'ParadaLinea'
type ParadaService interface {
	ParadasPorLinea(l *Linea) ([]*ParadaLinea, error)
	ParadasPorEmpresa(e *Empresa) ([]*Parada, error)
}

// LineaService has actions to fetch 'Linea' data from Cuando Llega City Bus API.
type LineaService interface {
	LineasPorEmpresa(e *Empresa) ([]*Linea, error)
}

// RecorridoService has actions to fetch 'Recorrido' data from Cuando Llega City Bus API.
type RecorridoService interface {
	RecorridoDeLinea(l *Linea) (*Recorrido, error)
	RecorridosPorEmpresa(e *Empresa) ([]*Recorrido, error)
}

// EmpresaService has actions to fetch 'Recorrido' data from Cuando Llega City Bus API.
type EmpresaService interface {
	ObtenerLineas(e *Empresa) error
}
