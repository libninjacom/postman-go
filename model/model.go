package model

type ScimUserResource struct {
	Schemas    []string    `json:"schemas"`
	Id         string      `json:"id"`
	UserName   string      `json:"userName"`
	Name       interface{} `json:"name"`
	ExternalId string      `json:"externalId"`
	Active     bool        `json:"active"`
	Meta       interface{} `json:"meta"`
}

type ImportExportFile struct {
	Type  string `json:"type"`
	Input string `json:"input"`
}

type JsonSchema struct {
	Type  string      `json:"type"`
	Input interface{} `json:"input"`
}

type JsonStringified struct {
	Type  string `json:"type"`
	Input string `json:"input"`
}
