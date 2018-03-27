package main

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

func main() {
	scli := swparadas.NewSWParadasSoap("", false, nil)

	in := &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       "WEB.SUR",
		Clave:         "PAR.SW.SUR",
		CodigoEmpresa: 355,
		IsSublinea:    false,
	}
	res, err := scli.RecuperarLineasPorCodigoEmpresa(in)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: '%+v'", res)

	var result swparadas.RecuperarLineasPorCodigoEmpresaResult
	json.Unmarshal([]byte(res.RecuperarLineasPorCodigoEmpresaResult), result)

	fmt.Println(result)
}
