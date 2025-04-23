package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/tabwriter"

	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/booscaaa/hamburguer-go/internal/core/dto"
	"github.com/spf13/viper"
)

type itemLLM struct{}

// GenerateRecomendation implements domain.ItemLLMGateway.
func (gateway *itemLLM) GenerateRecomendation(
	ctx context.Context,
	tools []dto.Tool,
	items []domain.Item,
) ([]dto.Function, error) {
	var b bytes.Buffer

	writer := tabwriter.NewWriter(&b, 0, 8, 1, '\t', tabwriter.AlignRight)
	for _, item := range items {
		fmt.Fprintf(writer, "Nome: %s\tPreço: %v\n", item.Name, item.Price)
	}
	writer.Flush()

	assistantPrompt := fmt.Sprintf(`
		Você é um atendente de restaurante.
		Sempre monte um cardápio para pedir seguindo essa lista de produtos: %s.
		Informe os itens que você pediu pedir para comer e o preço total da compra.
		Não dê mais informações do que o necessário.
		Sempre seja gentil.
		Seja sucinto!
		Diversifique sempre os pedidos para não ser toda vez a mesma coisa.
		Coloque uma pequena frase legal no final.
		Escreva os números por extenso sempre.
		Escreva tudo por extenso para leitura da Alexa.
		Escreva tudo em uma frase, respeitando o portugues.
		Não coloque caracteres especiais.
		Fale sempre em primeira pessoa.
		Use os tools sempre.
		Sempre retorne bebidas para o número de pessoas conseguir tomar sem faltar.
		Sempre retorne o valor total da compra por extenso.
		Sempre retorne uma frase legal no final informando os produtos junto.
		Sempre retorne tudo no mesmo array de itens.
		Fale como você já tivesse pedido os itens.
	`, b.String())

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
				"content": "Peça items para 2 pessoas jantar retornando os itens nos 2 tools",
			},
		},
		"tools":       dto.ToolsToOpenAi(tools),
		"tool_choice": "auto",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %v", err)
	}

	// fmt.Println(string(jsonPayload))

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response dto.Response

	if err = json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("no choices in response")
	}

	functions, err := response.GetFunctions()
	if err != nil {
		return nil, fmt.Errorf("error getting functions: %v", err)
	}

	return functions, nil
}

func NewItemLLM() domain.ItemLLMGateway {
	return &itemLLM{}
}
