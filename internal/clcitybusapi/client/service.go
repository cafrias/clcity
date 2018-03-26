package client

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

var _ clcitybusapi.Client = &Client{}

// Client represents a client to the 'Cuando Llega City Bus' API.
type Client struct {
	paradaService ParadaService
	lineaService  LineaService
}

// ParadaService returns an initialized instance of ParadaService.
func (c *Client) ParadaService() clcitybusapi.ParadaService { return &c.paradaService }

// LineaService returns an initialized instance of LineaService.
func (c *Client) LineaService() clcitybusapi.LineaService { return &c.lineaService }
