package client

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.Client = &Client{}

// Credentials to access Cuando Llega City Bus API.
const (
	Usuario = "WEB.SUR"
	Clave   = "PAR.SW.SUR"
)

// Client represents a client to the 'Cuando Llega City Bus' API.
type Client struct {
	client           SOAPClient
	paradaService    clcitybusapi.ParadaService
	lineaService     clcitybusapi.LineaService
	recorridoService clcitybusapi.RecorridoService
	empresaService   clcitybusapi.EmpresaService
}

// ParadaService returns an initialized instance of ParadaService.
func (c *Client) ParadaService() clcitybusapi.ParadaService { return c.paradaService }

// LineaService returns an initialized instance of LineaService.
func (c *Client) LineaService() clcitybusapi.LineaService { return c.lineaService }

// RecorridoService returns an initialized instance of RecorridoService.
func (c *Client) RecorridoService() clcitybusapi.RecorridoService { return c.recorridoService }

// EmpresaService returns an initialized instance of EmpresaService.
func (c *Client) EmpresaService() clcitybusapi.EmpresaService { return c.empresaService }

// SOAPClient represents the `cuando llega City Bus` API endpoint.
type SOAPClient interface {
	RecuperarLineasPorCodigoEmpresa(request *swparadas.RecuperarLineasPorCodigoEmpresa) (*swparadas.RecuperarLineasPorCodigoEmpresaResponse, error)
	RecuperarParadasCompletoPorLinea(request *swparadas.RecuperarParadasCompletoPorLinea) (*swparadas.RecuperarParadasCompletoPorLineaResponse, error)
	RecuperarRecorridoParaMapaPorEntidadYLinea(request *swparadas.RecuperarRecorridoParaMapaPorEntidadYLinea) (*swparadas.RecuperarRecorridoParaMapaPorEntidadYLineaResponse, error)
}

// NewClient creates a new client for communicating with `Cuando Llega City Bus` API.
func NewClient(scli SOAPClient, dumpPath string) *Client {
	if scli == nil {
		scli = swparadas.NewSWParadasSoap("", false, nil)
	}
	if dumpPath == "" {
		dumpPath = "."
	}

	lService := &LineaService{
		client: scli,
		Path:   dumpPath,
	}
	pService := &ParadaService{
		client:       scli,
		lineaService: lService,
		Path:         dumpPath,
	}
	rService := &RecorridoService{
		client:       scli,
		lineaService: lService,
		Path:         dumpPath,
	}
	return &Client{
		client:           scli,
		lineaService:     lService,
		paradaService:    pService,
		recorridoService: rService,
	}
}
