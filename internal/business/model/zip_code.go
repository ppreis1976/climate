package model

type ZipCode struct {
	Cep         ZipCodeID `json:"cep"`
	Logradouro  string    `json:"logradouro"`
	Complemento string    `json:"complemento"`
	Unidade     string    `json:"unidade"`
	Bairro      string    `json:"bairro"`
	Localidade  string    `json:"localidade"`
	Uf          string    `json:"uf"`
	Estado      string    `json:"estado"`
	Regiao      string    `json:"regiao"`
	Ibge        string    `json:"ibge"`
	Gia         string    `json:"gia"`
	Ddd         string    `json:"ddd"`
	Siafi       string    `json:"siafi"`
}
