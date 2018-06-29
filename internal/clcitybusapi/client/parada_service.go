package client

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/cafrias/clcity/internal/clcitybusapi/dump"

	"github.com/cafrias/clcity/pkg/geo"

	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.ParadaService = &ParadaService{}

// ParadaService has actions to fetch 'ParadaLinea' data from Cuando Llega City Bus API.
type ParadaService struct {
	scli SOAPClient
	cli  clcitybusapi.Client
	Path string
}

func (s *ParadaService) mapParadaLineaFromSW(swp *swparadas.ParadaLinea) (*clcitybusapi.ParadaLinea, error) {
	cod, err := strconv.Atoi(swp.Codigo)
	if err != nil {
		return nil, err
	}

	latStr := strings.Replace(swp.Latitud, ",", ".", -1)
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, err
	}

	longStr := strings.Replace(swp.Longitud, ",", ".", -1)
	Lon, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		return nil, err
	}

	callePrincipal := strings.Title(strings.ToLower(swp.CallePrincipal))
	calleInterseccion := strings.Title(strings.ToLower(swp.CalleInterseccion))

	return &clcitybusapi.ParadaLinea{
		Codigo:        cod,
		Nombre:        fmt.Sprintf("%s y %s", callePrincipal, calleInterseccion),
		Identificador: swp.Identificador,
		Descripcion:   swp.Descripcion,
		Punto: geo.Point{
			Lat: lat,
			Lon: Lon,
		},
	}, nil
}

func (s *ParadaService) fetchParadasPorLinea(linea *clcitybusapi.Linea, ret *[]*clcitybusapi.ParadaLinea) error {
	in := &swparadas.RecuperarParadasPorLineaParaCuandoLlega{
		Usuario:           Usuario,
		Clave:             Clave,
		CodigoLineaParada: int32(linea.Codigo),
		IsSubLinea:        false,
		IsInteligente:     false,
	}
	res, err := s.scli.RecuperarParadasPorLineaParaCuandoLlega(in)
	if err != nil {
		return err
	}

	result := new(swparadas.RecuperarParadasPorLineaParaCuandoLlegaResult)
	err = json.Unmarshal([]byte(res.RecuperarParadasPorLineaParaCuandoLlegaResult), result)
	if err != nil {
		return err
	}

	if result.CodigoEstado != 0 {
		fmt.Printf("No stops found for line id '%v', skipping ...\n", linea.Codigo)
	}

	for _, parada := range result.Paradas {
		p, err := s.mapParadaLineaFromSW(parada)
		if err != nil {
			return err
		}
		p.Linea = linea
		*ret = append(*ret, p)
	}

	return nil
}

// ParadasPorLinea fetches all 'ParadaLinea' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *ParadaService) ParadasPorLinea(linea *clcitybusapi.Linea) ([]*clcitybusapi.ParadaLinea, error) {
	outFile := fmt.Sprintf("%s/paradas_linea_%v.json", s.Path, linea.Codigo)

	var ret []*clcitybusapi.ParadaLinea

	// If dump
	if ok := dump.Read(&ret, outFile); ok == true {
		// Map linea
		for _, lp := range ret {
			lp.Linea = linea
		}
	} else {
		// If no dump
		err := s.fetchParadasPorLinea(linea, &ret)
		if err != nil {
			return nil, err
		}

		// Write dump file
		err = dump.Write(ret, outFile)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// ParadasPorEmpresa fetches all 'ParadaLinea' entities associated with given 'Empresa' identified by `CodigoEmpresa`.
func (s *ParadaService) ParadasPorEmpresa(empresa *clcitybusapi.Empresa) ([]*clcitybusapi.Parada, error) {
	outFile := fmt.Sprintf("%s/paradas_empresa.json", s.Path)

	var ret []*clcitybusapi.Parada
	if ok := dump.Read(&ret, outFile); ok == true {
		return ret, nil
	}

	lineas, err := s.cli.LineaService().LineasPorEmpresa(empresa)
	if err != nil {
		return nil, err
	}

	type RequestResult struct {
		Value []*clcitybusapi.ParadaLinea
		Error error
	}

	var wg sync.WaitGroup

	lineasQty := len(lineas)
	wg.Add(lineasQty)

	pStream := make(chan *RequestResult, lineasQty)
	for _, linea := range lineas {
		go func(linea *clcitybusapi.Linea) {
			defer wg.Done()
			result := new(RequestResult)
			res, err := s.ParadasPorLinea(linea)
			if err != nil {
				result.Error = err
				pStream <- result
				return
			}
			result.Value = res
			pStream <- result
			return
		}(linea)
	}

	wg.Wait()

	// Dedupe
	deduped := make(map[string]*clcitybusapi.ParadaLinea)
	for i := 0; i < lineasQty; i++ {
		result := <-pStream
		if result.Error != nil {
			return nil, result.Error
		}
		for _, linea := range result.Value {
			deduped[linea.Identificador] = linea
		}
	}

	// Map to correct type
	for _, linea := range deduped {
		retLin := &clcitybusapi.Parada{
			Codigo: linea.Identificador,
			Punto:  linea.Punto,
			Nombre: linea.Nombre,
		}
		ret = append(ret, retLin)
	}

	// Sort
	sort.Slice(ret, func(i, j int) bool { return ret[i].Codigo < ret[j].Codigo })

	// Write dump file
	err = dump.Write(ret, outFile)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
