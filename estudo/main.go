package main

import (
	"encoding/json"
	"fmt"

	Veiculo "github.com/Fernandao-lssilva/golang-estudo/estudo/veiculo"
)

func main() {
	astra := Veiculo.Veiculo{
		AnoFabricao: 2005,
		AnoModelo:   2006,
		Fabricante:  "GM",
		Modelo:      "Astra",
		Valor:       50000.0,
		Placa:       "qfl2345",
		Cor:         "Branco",
		Chassi:      "123456789",
	}
	//	fmt.Println("Bom dia, Rico")
	//entidades.ImprimeAnos()
	//fmt.Println(fmt.Sprintf("ano Fabricacao: %d", anoFabricacao))
	//fmt.Println(fmt.Sprintf("ano Modelo: %d", AnoModelo))

	bytes, err := json.MarshalIndent(astra, "", "  ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))
}
