package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"text/tabwriter"

	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/spf13/viper"
)

type reviewLLM struct{}

// GetTop3Reviews implements domain.ReviewLLMGateway.
func (r *reviewLLM) GetTop3Reviews(ctx context.Context, reviews []domain.Review) (*string, error) {
	var b bytes.Buffer

	writer := tabwriter.NewWriter(&b, 0, 8, 1, '\t', tabwriter.AlignRight)
	for _, review := range reviews {
		fmt.Fprintf(writer, "Nome: %s\tDescrição: %v\n", review.Name, review.Description)
	}
	writer.Flush()

	assistantPrompt := `
	Você é um assistente de inteligência artificial animado e divertido, ajudando na hora de premiar os destaques de uma palestra!
	
	Você receberá várias respostas de participantes sobre o evento. Sua missão é simples:
	- Filtrar só as respostas positivas — aquelas que mostram que a galera curtiu mesmo!
	- Escolher as 3 melhores avaliações positivas.
	- Falar apenas o nome dos 3 vencedores, sem explicações nem justificativas, porque aqui a emoção fala mais alto!
	- Ler em voz alta a melhor resposta entre as três, com destaque e entre aspas duplas.
	
	Importante: não mencione que as respostas foram filtradas ou que são apenas positivas. Apenas celebre os nomes e leia a melhor como se fosse natural!
	
	Ah, e finalize convidando os três campeões a subirem ao palco para receber um brinde especial do nosso querido palestrante Bosca
	
	Tudo em texto plano, estilo fala de palco, para a Alexa narrar em primeira pessoa com aquele entusiasmo de premiação!
	`

	client := &http.Client{}

	payload := map[string]any{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{
				"role":    "assistant",
				"content": assistantPrompt,
			},
			{
				"role":    "user",
				"content": fmt.Sprintf("Retorne as analises de acordo com as respostas %s", b.String()),
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("openai.api_key")))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var response map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	fmt.Println(response)

	choices, ok := response["choices"].([]any)
	if !ok || len(choices) == 0 {
		return nil, fmt.Errorf("no choices found in response")
	}
	firstChoice := choices[0].(map[string]any)
	message, ok := firstChoice["message"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("no message found in response")
	}
	content, ok := message["content"].(string)
	if !ok {
		return nil, fmt.Errorf("no content found in response")
	}
	return &content, nil
}

func NewReviewLLM() domain.ReviewLLMGateway {
	return &reviewLLM{}
}
