package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
APIEndpoint define la URL base de la API de análisis TLS de SSL Labs.

Documentación oficial:
https://github.com/ssllabs/ssllabs-scan/blob/master/ssllabs-api-docs.md
*/
const APIEndpoint = "https://api.ssllabs.com/api/v3/analyze"

/*
AnalysisResponse representa la respuesta principal
devuelta por la API de SSL Labs.
*/
type AnalysisResponse struct {
	Host      string     `json:"host"`
	Status    string     `json:"status"`
	Endpoints []Endpoint `json:"endpoints"`
}

/*
Endpoint representa un servidor individual (IP)
asociado al dominio analizado.
*/
type Endpoint struct {
	IPAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
}

// Funciones

/*
requestAnalysis consulta la API de SSL Labs y devuelve
el estado actual del análisis TLS para un dominio.
*/
func requestAnalysis(domain string) (*AnalysisResponse, error) {
	url := fmt.Sprintf(
		"%s?host=%s&startNew=off&fromCache=on",
		APIEndpoint,
		domain,
	)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var analysis AnalysisResponse
	if err := json.Unmarshal(body, &analysis); err != nil {
		return nil, err
	}

	return &analysis, nil
}

/*
readDomainFromConsole solicita al usuario el dominio
a analizar mediante la entrada estándar.
*/
func readDomainFromConsole() (string, error) {
	fmt.Print("Ingrese el dominio a analizar: ")

	reader := bufio.NewReader(os.Stdin)
	domain, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	domain = strings.TrimSpace(domain)
	if domain == "" {
		return "", fmt.Errorf("el dominio no puede estar vacío")
	}

	return domain, nil
}

// Programa principal

func main() {
	domain, err := readDomainFromConsole()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("\nIniciando análisis TLS para: %s\n\n", domain)

	for {
		analysis, err := requestAnalysis(domain)
		if err != nil {
			fmt.Println("Error consultando SSL Labs:", err)
			os.Exit(1)
		}

		fmt.Println("Estado actual:", analysis.Status)

		if analysis.Status == "READY" {
			fmt.Println("\nResultados del análisis:")

			for _, endpoint := range analysis.Endpoints {
				fmt.Printf(
					"- IP: %s | Grado TLS: %s\n",
					endpoint.IPAddress,
					endpoint.Grade,
				)
			}
			return
		}

		time.Sleep(15 * time.Second)
	}
}
