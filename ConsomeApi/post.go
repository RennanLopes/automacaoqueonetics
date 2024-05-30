package consomeapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func Post(url string, dados interface{}) int {
	jsonDados, err := json.Marshal(dados)
	if err != nil {
		log.Fatalf("erro ao serializar json: %v", err)
	}

	// Verificar o JSON que está sendo enviado
	log.Printf("POST %s\n%s\n", url, jsonDados)

	// Criar uma solicitação POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonDados))
	if err != nil {
		log.Fatalf("Erro ao criar solicitação: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Adiciona autenticação básica
	username := os.Getenv("API_USERNAME")
	password := os.Getenv("API_PASSWORD")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))

	// Enviar a solicitação
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erro ao fazer solicitação POST: %v", err)
	}
	defer resp.Body.Close()

	// Ler a resposta da requisição
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Falha na requisição. Status: %d: %s", resp.StatusCode, string(body))
		return 0
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Erro ao deserializar resposta em JSON: %v", err)
	}

	id, ok := result["id"].(float64)
	if !ok {
		log.Fatalf("Erro: id não encontrado ou não é um número")
	}
	name, ok := result["name"].(string)
	if !ok {
		log.Fatalf("Erro: name não encontrado ou não é uma string")
	}

	log.Printf("O id de Retorno é este: %v, nome: %v", id, name)
	return int(id)

}
