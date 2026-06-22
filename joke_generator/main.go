package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Joke representa uma piada da API
type Joke struct {
	ID    int    `json:"id"`
	Joke  string `json:"joke"`
	Type  string `json:"type"`
}

// JokeResponse representa a resposta da API de piadas
type JokeResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Joke    string `json:"joke"`
	Type    string `json:"type"`
}

// JokeGenerator gerencia requisições de piadas
type JokeGenerator struct {
	apiURL    string
	httpClient *http.Client
}

// NewJokeGenerator cria uma nova instância de JokeGenerator
func NewJokeGenerator() *JokeGenerator {
	return &JokeGenerator{
		apiURL: "https://official-joke-api.appspot.com",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetRandomJoke retorna uma piada aleatória
func (jg *JokeGenerator) GetRandomJoke() (string, error) {
	url := fmt.Sprintf("%s/random_joke", jg.apiURL)

	resp, err := jg.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("erro ao conectar à API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro na API: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %w", err)
	}

	var joke Joke
	if err := json.Unmarshal(body, &joke); err != nil {
		return "", fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return joke.Joke, nil
}

// GetRandomJokeByType retorna uma piada aleatória de um tipo específico
func (jg *JokeGenerator) GetRandomJokeByType(jokeType string) (string, error) {
	url := fmt.Sprintf("%s/jokes/%s/random", jg.apiURL, jokeType)

	resp, err := jg.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("erro ao conectar à API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("tipo de piada não encontrado: %s", jokeType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %w", err)
	}

	var joke Joke
	if err := json.Unmarshal(body, &joke); err != nil {
		return "", fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return joke.Joke, nil
}

// GetTenJokes retorna 10 piadas aleatórias
func (jg *JokeGenerator) GetTenJokes() ([]string, error) {
	url := fmt.Sprintf("%s/jokes/ten", jg.apiURL)

	resp, err := jg.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar à API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na API: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	var jokes []Joke
	if err := json.Unmarshal(body, &jokes); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	var jokeStrings []string
	for _, j := range jokes {
		jokeStrings = append(jokeStrings, j.Joke)
	}

	return jokeStrings, nil
}

func main() {
	fmt.Println("🎭 Random Joke Generator - Bem-vindo!\n")

	generator := NewJokeGenerator()

	// Exemplo 1: Uma piada aleatória
	fmt.Println("📝 Gerando uma piada aleatória...")
	joke, err := generator.GetRandomJoke()
	if err != nil {
		log.Printf("❌ Erro: %v\n", err)
	} else {
		fmt.Printf("😂 %s\n\n", joke)
	}

	// Exemplo 2: Piada por tipo (general)
	fmt.Println("📝 Gerando uma piada do tipo 'general'...")
	joke, err = generator.GetRandomJokeByType("general")
	if err != nil {
		log.Printf("❌ Erro: %v\n", err)
	} else {
		fmt.Printf("😂 %s\n\n", joke)
	}

	// Exemplo 3: Piada por tipo (knock-knock)
	fmt.Println("📝 Gerando uma piada do tipo 'knock-knock'...")
	joke, err = generator.GetRandomJokeByType("knock-knock")
	if err != nil {
		log.Printf("❌ Erro: %v\n", err)
	} else {
		fmt.Printf("😂 %s\n\n", joke)
	}

	// Exemplo 4: 10 piadas
	fmt.Println("📝 Gerando 10 piadas...")
	jokes, err := generator.GetTenJokes()
	if err != nil {
		log.Printf("❌ Erro: %v\n", err)
	} else {
		fmt.Printf("✅ %d piadas geradas:\n", len(jokes))
		for i, j := range jokes {
			fmt.Printf("%d. 😂 %s\n", i+1, j)
		}
	}
}
