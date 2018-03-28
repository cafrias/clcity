package client

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.Client = &Client{}

// Client represents a client to the 'Cuando Llega City Bus' API.
type Client struct {
	client        SOAPClient
	paradaService clcitybusapi.ParadaService
	lineaService  clcitybusapi.LineaService
}

// ParadaService returns an initialized instance of ParadaService.
func (c *Client) ParadaService() clcitybusapi.ParadaService { return c.paradaService }

// LineaService returns an initialized instance of LineaService.
func (c *Client) LineaService() clcitybusapi.LineaService { return c.lineaService }

// SOAPClient represents the `cuando llega City Bus` API endpoint.
type SOAPClient interface {
	RecuperarLineasPorCodigoEmpresa(request *swparadas.RecuperarLineasPorCodigoEmpresa) (*swparadas.RecuperarLineasPorCodigoEmpresaResponse, error)
	RecuperarParadasCompletoPorLinea(request *swparadas.RecuperarParadasCompletoPorLinea) (*swparadas.RecuperarParadasCompletoPorLineaResponse, error)
}

// NewClient creates a new client for communicating with `Cuando Llega City Bus` API.
func NewClient(scli SOAPClient) *Client {
	if scli == nil {
		scli = swparadas.NewSWParadasSoap("", false, nil)
	}

	lService := &LineaService{
		client: scli,
	}
	pService := &ParadaService{
		client:       scli,
		lineaService: lService,
	}
	return &Client{
		client:        scli,
		lineaService:  lService,
		paradaService: pService,
	}
}
