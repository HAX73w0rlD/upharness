package core

// ModelInfo describes an available LLM model.
type ModelInfo struct {
	ID       string
	Provider string
	Free     bool
}

// Catalog returns a minimal model catalog.
func Catalog() []ModelInfo {
	return []ModelInfo{
		{ID: "z-ai/glm-5.3-flash", Provider: "nvidia", Free: true},
		{ID: "meta/llama-3.1-8b-instruct", Provider: "openrouter", Free: true},
	}
}
