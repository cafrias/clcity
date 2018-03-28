package main

import (
	"fmt"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
)

func main() {
	cli := client.NewClient()
	cli.Connect(nil)
	ret, err := cli.LineaService().LineasPorEmpresa(355)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: '%#v'", ret)
}
