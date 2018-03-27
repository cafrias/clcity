package client

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.Client = &Client{}

// Client represents a client to the 'Cuando Llega City Bus' API.
type Client struct {
	soap          *swparadas.SWParadasSoap
	paradaService ParadaService
	lineaService  LineaService
}

// ParadaService returns an initialized instance of ParadaService.
func (c *Client) ParadaService() clcitybusapi.ParadaService { return &c.paradaService }

// LineaService returns an initialized instance of LineaService.
func (c *Client) LineaService() clcitybusapi.LineaService { return &c.lineaService }

// NewClient creates a new client for communicating with `Cuando Llega City Bus` API.
func NewClient(scli *swparadas.SWParadasSoap) *Client {
	return &Client{
		soap: scli,
		paradaService: ParadaService{
			client: scli,
		},
		lineaService: LineaService{
			client: scli,
		},
	}
}
