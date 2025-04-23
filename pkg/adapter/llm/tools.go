package llm

import (
	"github.com/booscaaa/hamburguer-go/internal/core/dto"
)

var Tools = []dto.Tool{}

func InitializeTools() {
	tool := dto.NewTool(
		"get_hamburger_items",
		"recomendation",
		`Obter a lista de nomes dos hamburgueres 
			disponíveis para pedido. Retorne apenas os nomes exatos 
			dos itens como aparecem no site, sem descrições adicionais, 
			para permitir a realização do pedido`,
		map[string]any{
			"type": "object",
			"properties": map[string]any{
				"items": map[string]any{
					"type":        "array",
					"description": "The list of hamburger items",
					"items": map[string]any{
						"type": "string",
					},
				},
			},
			"required": []string{"items"},
		},
	)

	toolResponse := dto.NewTool(
		"get_alexa_response",
		"recomendation",
		`Gere uma resposta amigável e envolvente para a Alexa dizer ao 
			confirmar ou reconhecer pedidos dos clientes. A resposta deve 
			ser natural, conversacional e apropriada para o contexto. Sempre use esse tool`,
		map[string]any{
			"type": "object",
			"properties": map[string]any{
				"response": map[string]any{
					"type":        "string",
					"description": "Uma resposta amigável para a Alexa dizer",
				},
			},
			"required": []string{"response"},
		},
	)

	Tools = append(Tools, toolResponse)

	Tools = append(Tools, tool)
}

func GetRecomendationTools() []dto.Tool {
	var tools []dto.Tool
	for _, tool := range Tools {
		if tool.Purpose == "recomendation" {
			tools = append(tools, tool)
		}
	}

	return tools
}
