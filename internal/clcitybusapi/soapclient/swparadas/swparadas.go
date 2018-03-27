package swparadas

import (
	"encoding/xml"
	"time"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type RecuperarLineaPorLocalidad struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarLineaPorLocalidad"`

	Usuario    string `xml:"usuario,omitempty"`
	Clave      string `xml:"clave,omitempty"`
	Localidad  string `xml:"localidad,omitempty"`
	Provincia  string `xml:"provincia,omitempty"`
	Pais       string `xml:"pais,omitempty"`
	IsSublinea bool   `xml:"isSublinea,omitempty"`
}

type RecuperarLineaPorLocalidadResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarLineaPorLocalidadResponse"`

	RecuperarLineaPorLocalidadResult string `xml:"RecuperarLineaPorLocalidadResult,omitempty"`
}

type RecuperarLineaPorEntidad struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarLineaPorEntidad"`

	Usuario          string `xml:"usuario,omitempty"`
	Clave            string `xml:"clave,omitempty"`
	CodigoEntidadSMP int32  `xml:"codigoEntidadSMP,omitempty"`
	IsSublinea       bool   `xml:"isSublinea,omitempty"`
}

type RecuperarLineaPorEntidadResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarLineaPorEntidadResponse"`

	RecuperarLineaPorEntidadResult string `xml:"RecuperarLineaPorEntidadResult,omitempty"`
}

type RecuperarLineaPorCuandoLlega struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarLineaPorCuandoLlega"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoCuandoLlega int32  `xml:"codigoCuandoLlega,omitempty"`
	IsSublinea        bool   `xml:"isSublinea,omitempty"`
}

type RecuperarLineaPorCuandoLlegaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarLineaPorCuandoLlegaResponse"`

	RecuperarLineaPorCuandoLlegaResult string `xml:"RecuperarLineaPorCuandoLlegaResult,omitempty"`
}

type RecuperarParadasCompletoPorLinea struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasCompletoPorLinea"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSublinea        bool   `xml:"isSublinea,omitempty"`
	IsInteligente     bool   `xml:"isInteligente,omitempty"`
}

type RecuperarParadasCompletoPorLineaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasCompletoPorLineaResponse"`

	RecuperarParadasCompletoPorLineaResult string `xml:"RecuperarParadasCompletoPorLineaResult,omitempty"`
}

type RecuperarParadasConBanderaYDestinoPorLinea struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasConBanderaYDestinoPorLinea"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSublinea        bool   `xml:"isSublinea,omitempty"`
	IsInteligente     bool   `xml:"isInteligente,omitempty"`
}

type RecuperarParadasConBanderaYDestinoPorLineaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasConBanderaYDestinoPorLineaResponse"`

	RecuperarParadasConBanderaYDestinoPorLineaResult string `xml:"RecuperarParadasConBanderaYDestinoPorLineaResult,omitempty"`
}

type RecuperarParadasPorLineaCalleEInterseccion struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasPorLineaCalleEInterseccion"`

	Usuario            string `xml:"usuario,omitempty"`
	Clave              string `xml:"clave,omitempty"`
	CodigoLinea        int32  `xml:"codigoLinea,omitempty"`
	CodigoCalle        int32  `xml:"codigoCalle,omitempty"`
	CodigoInterseccion int32  `xml:"codigoInterseccion,omitempty"`
	IsSublinea         bool   `xml:"isSublinea,omitempty"`
	IsInteligente      bool   `xml:"isInteligente,omitempty"`
}

type RecuperarParadasPorLineaCalleEInterseccionResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasPorLineaCalleEInterseccionResponse"`

	RecuperarParadasPorLineaCalleEInterseccionResult string `xml:"RecuperarParadasPorLineaCalleEInterseccionResult,omitempty"`
}

type RecuperarParadasConBanderaPorLineaCalleEInterseccion struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasConBanderaPorLineaCalleEInterseccion"`

	Usuario            string `xml:"usuario,omitempty"`
	Clave              string `xml:"clave,omitempty"`
	CodigoLineaParada  int32  `xml:"codigoLineaParada,omitempty"`
	CodigoCalle        int32  `xml:"codigoCalle,omitempty"`
	CodigoInterseccion int32  `xml:"codigoInterseccion,omitempty"`
	IsSubLinea         bool   `xml:"isSubLinea,omitempty"`
	IsInteligente      bool   `xml:"isInteligente,omitempty"`
}

type RecuperarParadasConBanderaPorLineaCalleEInterseccionResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasConBanderaPorLineaCalleEInterseccionResponse"`

	RecuperarParadasConBanderaPorLineaCalleEInterseccionResult string `xml:"RecuperarParadasConBanderaPorLineaCalleEInterseccionResult,omitempty"`
}

type RecuperarParadasPorLineaParaCuandoLlega struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasPorLineaParaCuandoLlega"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSubLinea        bool   `xml:"isSubLinea,omitempty"`
	IsInteligente     bool   `xml:"isInteligente,omitempty"`
}

type RecuperarParadasPorLineaParaCuandoLlegaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasPorLineaParaCuandoLlegaResponse"`

	RecuperarParadasPorLineaParaCuandoLlegaResult string `xml:"RecuperarParadasPorLineaParaCuandoLlegaResult,omitempty"`
}

type RecuperarParadasMasCercanasPorLocalidadProvinciaPais struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasMasCercanasPorLocalidadProvinciaPais"`

	Usuario              string  `xml:"usuario,omitempty"`
	Clave                string  `xml:"clave,omitempty"`
	Latitud              float64 `xml:"latitud,omitempty"`
	Longitud             float64 `xml:"longitud,omitempty"`
	ListaCodigosEmpresa  string  `xml:"listaCodigosEmpresa,omitempty"`
	DescripcionProvincia string  `xml:"descripcionProvincia,omitempty"`
	DescripcionPais      string  `xml:"descripcionPais,omitempty"`
	IsInteligente        bool    `xml:"isInteligente,omitempty"`
}

type RecuperarParadasMasCercanasPorLocalidadProvinciaPaisResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasMasCercanasPorLocalidadProvinciaPaisResponse"`

	RecuperarParadasMasCercanasPorLocalidadProvinciaPaisResult string `xml:"RecuperarParadasMasCercanasPorLocalidadProvinciaPaisResult,omitempty"`
}

type RecuperarCallesPrincipalPorLinea struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarCallesPrincipalPorLinea"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSubLinea        bool   `xml:"isSubLinea,omitempty"`
}

type RecuperarCallesPrincipalPorLineaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarCallesPrincipalPorLineaResponse"`

	RecuperarCallesPrincipalPorLineaResult string `xml:"RecuperarCallesPrincipalPorLineaResult,omitempty"`
}

type RecuperarInterseccionPorLineaYCalle struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarInterseccionPorLineaYCalle"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	CodigoCalle       int32  `xml:"codigoCalle,omitempty"`
	IsSubLinea        bool   `xml:"isSubLinea,omitempty"`
}

type RecuperarInterseccionPorLineaYCalleResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarInterseccionPorLineaYCalleResponse"`

	RecuperarInterseccionPorLineaYCalleResult string `xml:"RecuperarInterseccionPorLineaYCalleResult,omitempty"`
}

type RecuperarProximosArribos struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarProximosArribos"`

	Usuario             string `xml:"usuario,omitempty"`
	Clave               string `xml:"clave,omitempty"`
	IdentificadorParada string `xml:"identificadorParada,omitempty"`
	CodigoLineaParada   int32  `xml:"codigoLineaParada,omitempty"`
	CodigoAplicacion    int32  `xml:"codigoAplicacion,omitempty"`
	Localidad           string `xml:"localidad,omitempty"`
	IsSublinea          bool   `xml:"isSublinea,omitempty"`
	IsSoloAdaptados     bool   `xml:"isSoloAdaptados,omitempty"`
}

type RecuperarProximosArribosResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarProximosArribosResponse"`

	RecuperarProximosArribosResult string `xml:"RecuperarProximosArribosResult,omitempty"`
}

type RecuperarProximosArribosSMS struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarProximosArribosSMS"`

	Usuario        string `xml:"usuario,omitempty"`
	Clave          string `xml:"clave,omitempty"`
	GatewayID      string `xml:"gatewayID,omitempty"`
	NumeroCelular  string `xml:"numeroCelular,omitempty"`
	NumeroOrigen   string `xml:"numeroOrigen,omitempty"`
	CodigoOperador int32  `xml:"codigoOperador,omitempty"`
	Mensaje        string `xml:"mensaje,omitempty"`
	IsSublinea     bool   `xml:"isSublinea,omitempty"`
}

type RecuperarProximosArribosSMSResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarProximosArribosSMSResponse"`

	RecuperarProximosArribosSMSResult string `xml:"RecuperarProximosArribosSMSResult,omitempty"`
}

type RecuperarTodasEmpresasAMigrar struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarTodasEmpresasAMigrar"`

	Usuario       string `xml:"usuario,omitempty"`
	Clave         string `xml:"clave,omitempty"`
	CodigoEntidad int32  `xml:"codigoEntidad,omitempty"`
	CodigoLinea   int32  `xml:"codigoLinea,omitempty"`
}

type RecuperarTodasEmpresasAMigrarResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarTodasEmpresasAMigrarResponse"`

	RecuperarTodasEmpresasAMigrarResult string `xml:"RecuperarTodasEmpresasAMigrarResult,omitempty"`
}

type RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLinea struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLinea"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSublinea        bool   `xml:"isSublinea,omitempty"`
}

type RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLineaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLineaResponse"`

	RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLineaResult string `xml:"RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLineaResult,omitempty"`
}

type RecuperarRecorridoParaMapaPorEntidadYLinea struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarRecorridoParaMapaPorEntidadYLinea"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSublinea        bool   `xml:"isSublinea,omitempty"`
}

type RecuperarRecorridoParaMapaPorEntidadYLineaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarRecorridoParaMapaPorEntidadYLineaResponse"`

	RecuperarRecorridoParaMapaPorEntidadYLineaResult string `xml:"RecuperarRecorridoParaMapaPorEntidadYLineaResult,omitempty"`
}

type RecuperarPuestosRecargasMasCercanosPorLocalidad struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarPuestosRecargasMasCercanosPorLocalidad"`

	Usuario            string  `xml:"usuario,omitempty"`
	Clave              string  `xml:"clave,omitempty"`
	Latitud            float64 `xml:"latitud,omitempty"`
	Longitud           float64 `xml:"longitud,omitempty"`
	ListaCodigoEmpresa string  `xml:"listaCodigoEmpresa,omitempty"`
}

type RecuperarPuestosRecargasMasCercanosPorLocalidadResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarPuestosRecargasMasCercanosPorLocalidadResponse"`

	RecuperarPuestosRecargasMasCercanosPorLocalidadResult string `xml:"RecuperarPuestosRecargasMasCercanosPorLocalidadResult,omitempty"`
}

type RecuperarPuntosDeRecargaPorLocalidad struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarPuntosDeRecargaPorLocalidad"`

	Usuario   string `xml:"usuario,omitempty"`
	Clave     string `xml:"clave,omitempty"`
	Localidad string `xml:"localidad,omitempty"`
	Provincia string `xml:"provincia,omitempty"`
	Pais      string `xml:"pais,omitempty"`
}

type RecuperarPuntosDeRecargaPorLocalidadResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarPuntosDeRecargaPorLocalidadResponse"`

	RecuperarPuntosDeRecargaPorLocalidadResult string `xml:"RecuperarPuntosDeRecargaPorLocalidadResult,omitempty"`
}

type RecuperarBanderasAsociadasAParada struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarBanderasAsociadasAParada"`

	Usuario             string `xml:"usuario,omitempty"`
	Clave               string `xml:"clave,omitempty"`
	IdentificadorParada string `xml:"identificadorParada,omitempty"`
}

type RecuperarBanderasAsociadasAParadaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarBanderasAsociadasAParadaResponse"`

	RecuperarBanderasAsociadasAParadaResult string `xml:"RecuperarBanderasAsociadasAParadaResult,omitempty"`
}

type RecuperarMensajes struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarMensajes"`

	Usuario             string `xml:"usuario,omitempty"`
	Clave               string `xml:"clave,omitempty"`
	ListaCodigosEmpresa string `xml:"listaCodigosEmpresa,omitempty"`
}

type RecuperarMensajesResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarMensajesResponse"`

	RecuperarMensajesResult string `xml:"RecuperarMensajesResult,omitempty"`
}

type SWParadasSoap struct {
	client SOAPClient
}

func NewSWParadasSoap(url string, tls bool, auth *soapclient.BasicAuth) *SWParadasSoap {
	if url == "" {
		url = "http://clsw.smartmovepro.net/ModuloParadas/SWParadas.asmx"
	}
	client := soapclient.NewSOAPClient(url, tls, auth)

	return &SWParadasSoap{
		client: client,
	}
}

func NewSWParadasSoapWithClient(client SOAPClient) *SWParadasSoap {
	return &SWParadasSoap{
		client: client,
	}
}

func (service *SWParadasSoap) RecuperarLineaPorLocalidad(request *RecuperarLineaPorLocalidad) (*RecuperarLineaPorLocalidadResponse, error) {
	response := new(RecuperarLineaPorLocalidadResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarLineaPorLocalidad", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarLineaPorEntidad(request *RecuperarLineaPorEntidad) (*RecuperarLineaPorEntidadResponse, error) {
	response := new(RecuperarLineaPorEntidadResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarLineaPorEntidad", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarLineasPorCodigoEmpresa(request *RecuperarLineasPorCodigoEmpresa) (*RecuperarLineasPorCodigoEmpresaResponse, error) {
	response := new(RecuperarLineasPorCodigoEmpresaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarLineasPorCodigoEmpresa", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarLineaPorCuandoLlega(request *RecuperarLineaPorCuandoLlega) (*RecuperarLineaPorCuandoLlegaResponse, error) {
	response := new(RecuperarLineaPorCuandoLlegaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarLineaPorCuandoLlega", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarParadasCompletoPorLinea(request *RecuperarParadasCompletoPorLinea) (*RecuperarParadasCompletoPorLineaResponse, error) {
	response := new(RecuperarParadasCompletoPorLineaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasCompletoPorLinea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarParadasConBanderaYDestinoPorLinea(request *RecuperarParadasConBanderaYDestinoPorLinea) (*RecuperarParadasConBanderaYDestinoPorLineaResponse, error) {
	response := new(RecuperarParadasConBanderaYDestinoPorLineaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasConBanderaYDestinoPorLinea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarParadasPorLineaCalleEInterseccion(request *RecuperarParadasPorLineaCalleEInterseccion) (*RecuperarParadasPorLineaCalleEInterseccionResponse, error) {
	response := new(RecuperarParadasPorLineaCalleEInterseccionResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasPorLineaCalleEInterseccion", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarParadasConBanderaPorLineaCalleEInterseccion(request *RecuperarParadasConBanderaPorLineaCalleEInterseccion) (*RecuperarParadasConBanderaPorLineaCalleEInterseccionResponse, error) {
	response := new(RecuperarParadasConBanderaPorLineaCalleEInterseccionResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasConBanderaPorLineaCalleEInterseccion", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarParadasPorLineaParaCuandoLlega(request *RecuperarParadasPorLineaParaCuandoLlega) (*RecuperarParadasPorLineaParaCuandoLlegaResponse, error) {
	response := new(RecuperarParadasPorLineaParaCuandoLlegaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasPorLineaParaCuandoLlega", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarParadasMasCercanasPorLocalidadProvinciaPais(request *RecuperarParadasMasCercanasPorLocalidadProvinciaPais) (*RecuperarParadasMasCercanasPorLocalidadProvinciaPaisResponse, error) {
	response := new(RecuperarParadasMasCercanasPorLocalidadProvinciaPaisResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasMasCercanasPorLocalidadProvinciaPais", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarCallesPrincipalPorLinea(request *RecuperarCallesPrincipalPorLinea) (*RecuperarCallesPrincipalPorLineaResponse, error) {
	response := new(RecuperarCallesPrincipalPorLineaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarCallesPrincipalPorLinea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarInterseccionPorLineaYCalle(request *RecuperarInterseccionPorLineaYCalle) (*RecuperarInterseccionPorLineaYCalleResponse, error) {
	response := new(RecuperarInterseccionPorLineaYCalleResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarInterseccionPorLineaYCalle", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarProximosArribos(request *RecuperarProximosArribos) (*RecuperarProximosArribosResponse, error) {
	response := new(RecuperarProximosArribosResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarProximosArribos", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarProximosArribosSMS(request *RecuperarProximosArribosSMS) (*RecuperarProximosArribosSMSResponse, error) {
	response := new(RecuperarProximosArribosSMSResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarProximosArribosSMS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarTodasEmpresasAMigrar(request *RecuperarTodasEmpresasAMigrar) (*RecuperarTodasEmpresasAMigrarResponse, error) {
	response := new(RecuperarTodasEmpresasAMigrarResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarTodasEmpresasAMigrar", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLinea(request *RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLinea) (*RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLineaResponse, error) {
	response := new(RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLineaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarRecorridoParaMapaAbrevYAmpliPorEntidadYLinea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarRecorridoParaMapaPorEntidadYLinea(request *RecuperarRecorridoParaMapaPorEntidadYLinea) (*RecuperarRecorridoParaMapaPorEntidadYLineaResponse, error) {
	response := new(RecuperarRecorridoParaMapaPorEntidadYLineaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarRecorridoParaMapaPorEntidadYLinea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarPuestosRecargasMasCercanosPorLocalidad(request *RecuperarPuestosRecargasMasCercanosPorLocalidad) (*RecuperarPuestosRecargasMasCercanosPorLocalidadResponse, error) {
	response := new(RecuperarPuestosRecargasMasCercanosPorLocalidadResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarPuestosRecargasMasCercanosPorLocalidad", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarPuntosDeRecargaPorLocalidad(request *RecuperarPuntosDeRecargaPorLocalidad) (*RecuperarPuntosDeRecargaPorLocalidadResponse, error) {
	response := new(RecuperarPuntosDeRecargaPorLocalidadResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarPuntosDeRecargaPorLocalidad", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarBanderasAsociadasAParada(request *RecuperarBanderasAsociadasAParada) (*RecuperarBanderasAsociadasAParadaResponse, error) {
	response := new(RecuperarBanderasAsociadasAParadaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarBanderasAsociadasAParada", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SWParadasSoap) RecuperarMensajes(request *RecuperarMensajes) (*RecuperarMensajesResponse, error) {
	response := new(RecuperarMensajesResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarMensajes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type SOAPClient interface {
	Call(soapAction string, request, response interface{}) error
}
