package core

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// ModelInfo beschreibt ein verfugbares LLM-Modell.
type ModelInfo struct {
	ID       string
	Provider string
	Free     bool
}

// Catalog liefert die verfugbaren Modelle.
func Catalog() []ModelInfo {
	return []ModelInfo{
		{ID: "nvidia/nemotron-3.5-lightning-30b-a3b", Provider: "nvidia", Free: true},
		{ID: "meta/llama-3.1-8b-instruct", Provider: "openrouter", Free: true},
	}
}

// LiveModels holt Modelle live von der NVIDIA API, wenn ein Key gesetzt ist.
func LiveModels() []ModelInfo {
	key := os.Getenv("NVIDIA_API_KEY")
	if key == "" { return Catalog() }
	client := &http.Client{Timeout: 20 * time.Second}
	req, err := http.NewRequest("GET", "https://integrate.api.nvidia.com/v1/models", nil)
	if err != nil { return Catalog() }
	req.Header.Set("Authorization", "Bearer "+key)
	resp, err := client.Do(req)
	if err != nil { return Catalog() }
	defer resp.Body.Close();
	if resp.StatusCode != 200 { return Catalog() }
	var body struct {
		Data []struct { ID string `json:"id"` } `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil { return Catalog() }
	out := make([]ModelInfo, 0, len(body.Data))
	for _, m := range body.Data {
		if m.ID != "" {
			out = append(out, ModelInfo{ID: m.ID, Provider: "nvidia", Free: true})
		}
	}
	return out;
}
