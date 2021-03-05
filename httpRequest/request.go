package httpRequest

import (
	"log"
	"net/http"
)

// GetRequest é uma função que faz um GET na página e retorna o statusCode da página. Caso não encontre a página, ocorrerá um Fatal error.
func GetRequest(url string) int {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("A página não foi encontrada")
	}
	return res.StatusCode
}
