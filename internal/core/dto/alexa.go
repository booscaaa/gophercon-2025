package dto

type AlexaRequest struct {
	Request struct {
		Type   string `json:"type"`
		Intent struct {
			Name  string          `json:"name"`
			Slots map[string]Slot `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
}

type Slot struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AlexaResponse struct {
	Version  string `json:"version"`
	Response struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
		ShouldEndSession bool `json:"shouldEndSession"`
	} `json:"response"`
}
