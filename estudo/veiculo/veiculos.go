package Veiculo

type Veiculo struct {
	AnoFabricao int     `json:"ano_fabricacao"`
	AnoModelo   int     `json:"ano-modelo"`
	Fabricante  string  `json:"fabricante"`
	Modelo      string  `json:"modelo"`
	Valor       float32 `json:"valor"`
	Placa       string  `json:"placa"`
	Cor         string  `json:"cor"`
	Chassi      string  `json:"chassi"`
}
