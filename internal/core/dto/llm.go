package dto

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Content   string     `json:"content"`
	ToolCalls []ToolCall `json:"tool_calls"`
}

type ToolCall struct {
	Function FunctionCall `json:"function"`
}

type FunctionCall struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
	Purpose  string   `json:"purpose"`
}

type Function struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
}

func NewTool(
	name string,
	purpose string,
	description string,
	parameters map[string]any,
) Tool {
	return Tool{
		Type:    "function",
		Purpose: purpose,
		Function: Function{
			Name:        name,
			Description: description,
			Parameters:  parameters,
		},
	}
}

func ToolsToOpenAi(tools []Tool) []map[string]any {
	var toolsOpenAi []map[string]any
	for _, tool := range tools {
		toolsOpenAi = append(toolsOpenAi, map[string]any{
			"type": tool.Type,
			"function": map[string]any{
				"name":        tool.Function.Name,
				"description": tool.Function.Description,
				"parameters":  tool.Function.Parameters,
			},
		})
	}

	return toolsOpenAi
}

func (response *Response) GetFunctions() ([]Function, error) {
	var functions []Function
	if len(response.Choices[0].Message.ToolCalls) > 0 {
		for _, toolCall := range response.Choices[0].Message.ToolCalls {
			var arg string
			var args map[string]any

			if err := json.Unmarshal(toolCall.Function.Arguments, &arg); err != nil {
				if err := json.Unmarshal(toolCall.Function.Arguments, &args); err != nil {
					return nil, fmt.Errorf("error parsing tool arguments: %v", err)
				}
			} else {
				if err := json.Unmarshal([]byte(arg), &args); err != nil {
					return nil, fmt.Errorf("error parsing tool arguments string: %v", err)
				}
			}

			functions = append(functions, Function{
				Name:       toolCall.Function.Name,
				Parameters: args,
			})
		}
	}
	return functions, nil
}
