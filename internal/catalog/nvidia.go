package catalog

import (
"encoding/json"
"fmt"
"net/http"
"os"
"time"
)

// NvidiaModels fetches the live model list from the NVIDIA API.
// Needs NVIDIA_API_KEY in the environment. Returns IDs like
// "nvidia/llama-3.1-nemotron-70b-instruct".
func NvidiaModels() ([]string, error) {
	key := os.Getenv("NVIDIA_API_KEY")
	if key == "" {
		return nil, fmt.Errorf("NVIDIA_API_KEY not set")
	}
	client := &http.Client{Timeout: 20 * time.Second}
	req, err := http.NewRequest("GET", "https://integrate.api.nvidia.com/v1/models", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+key)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("nvidia api status %d", resp.StatusCode)
	}
	var body struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(body.Data))
	for _, m := range body.Data {
		if m.ID != "" {
			ids = append(ids, m.ID)
		}
	}
	return ids, nil
}
