package service

type TranslationService struct{}

type VSCodeCopilotMCPConfig struct {
	Servers map[string]Server `json:"servers"`
	Inputs  []Input           `json:"inputs"`
}

type Server struct {
	Command string            `json:"command"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`
}

type Input struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Password    bool   `json:"password"`
	Description string `json:"description"`
}

type Skill struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type A2ACard struct {
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	URL                string          `json:"url"`
	Capabilities       map[string]bool `json:"capabilities"`
	DefaultInputModes  []string        `json:"defaultInputModes"`
	DefaultOutputModes []string        `json:"defaultOutputModes"`
	Skills             []Skill         `json:"skills"`
}
