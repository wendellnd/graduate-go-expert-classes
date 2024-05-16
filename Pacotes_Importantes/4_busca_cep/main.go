package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		request, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Printf("request error: %s \n", err.Error())
		}
		defer request.Body.Close()

		res, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Printf("error reading body: %s \n", err.Error())
		}

		var data ViaCepResponse
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Printf("error unmarshalling response: %s \n", err.Error())
		}
		fmt.Printf("ViaCEP response: %+v", data)

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Printf("error creating file: %s \n", err.Error())
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
		if err != nil {
			fmt.Printf("error writing file: %s", err.Error())
		}
	}
}
