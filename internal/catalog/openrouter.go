package catalog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func OpenRouterModels() ([]string, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Get("https://openrouter.ai/api/v1/models")
	if err != nil { return nil, err }
	defer resp.Body.Close()
	if resp.StatusCode != 200 { return nil, fmt.Errorf("OpenRouter API Status %d", resp.StatusCode) }
	var body struct {
		Data []struct { ID string `json:"id"` } `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil { return nil, err }
	ids := make([]string, 0, len(body.Data))
	for _, m := range body.Data { if m.ID != "" { ids = append(ids, m.ID) } }
	return ids, nil
}
