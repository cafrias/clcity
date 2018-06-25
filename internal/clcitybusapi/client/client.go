package client

import (
	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/internal/clcitybusapi/soapclient/swparadas"
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
}

// ParadaService returns an initialized instance of ParadaService.
func (c *Client) ParadaService() clcitybusapi.ParadaService { return c.paradaService }

// LineaService returns an initialized instance of LineaService.
func (c *Client) LineaService() clcitybusapi.LineaService { return c.lineaService }

// RecorridoService returns an initialized instance of RecorridoService.
func (c *Client) RecorridoService() clcitybusapi.RecorridoService { return c.recorridoService }

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

	cli := new(Client)

	rService := &RecorridoService{
		scli: scli,
		Path: dumpPath,
	}
	lService := &LineaService{
		scli: scli,
		cli:  cli,
		Path: dumpPath,
	}
	pService := &ParadaService{
		scli: scli,
		cli:  cli,
		Path: dumpPath,
	}

	cli.client = scli
	cli.lineaService = lService
	cli.paradaService = pService
	cli.recorridoService = rService

	return cli
}

// SetSOAPClient sets the SOAP client
func (c *Client) SetSOAPClient(scli SOAPClient) {
	c.client = scli
}

// SetRecorridoService sets RecorridoService.
func (c *Client) SetRecorridoService(s clcitybusapi.RecorridoService) {
	c.recorridoService = s
}

// SetParadaService sets ParadaService.
func (c *Client) SetParadaService(s clcitybusapi.ParadaService) {
	c.paradaService = s
}

// SetLineaService sets LineaService.
func (c *Client) SetLineaService(s clcitybusapi.LineaService) {
	c.lineaService = s
}
