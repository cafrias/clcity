package client

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.Client = &Client{}

// Client represents a client to the 'Cuando Llega City Bus' API.
type Client struct {
	client        SOAPClient
	paradaService ParadaService
	lineaService  LineaService
}

// ParadaService returns an initialized instance of ParadaService.
func (c *Client) ParadaService() clcitybusapi.ParadaService { return &c.paradaService }

// LineaService returns an initialized instance of LineaService.
func (c *Client) LineaService() clcitybusapi.LineaService { return &c.lineaService }

// Connect instantiates a default SOAP client.
func (c *Client) Connect(scli SOAPClient) {
	if scli == nil {
		scli = swparadas.NewSWParadasSoap("", false, nil)
	}
	c.client = scli
	c.paradaService.client = scli
	c.lineaService.client = scli
}

// SOAPClient represents the `cuando llega City Bus` API endpoint.
type SOAPClient interface {
	RecuperarLineasPorCodigoEmpresa(request *swparadas.RecuperarLineasPorCodigoEmpresa) (*swparadas.RecuperarLineasPorCodigoEmpresaResponse, error)
}

// NewClient creates a new client for communicating with `Cuando Llega City Bus` API.
func NewClient() *Client {
	return &Client{
		paradaService: ParadaService{},
		lineaService:  LineaService{},
	}
}
