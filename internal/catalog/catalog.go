package catalog

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type ModelEntry struct {
	ID        string
	Provider  string
	Unhealthy bool
	LastCheck int64
	LatencyMs int
}

type Provider struct {
	ID  string
	Free bool
}

func Providers() []Provider {
	return []Provider{{ID: "nvidia", Free: true}, {ID: "opencode", Free: true}}
}

func StaticFallback() []ModelEntry {
	return []ModelEntry{
		{ID: "nvidia/nemotron-3.5-lightning-30b-a3b", Provider: "nvidia"},
		{ID: "meta/llama-3.1-8b-instruct", Provider: "openrouter"},
	}
}

func capFile() string {
	d := os.Getenv("HOME")
	if d == "" { d = "." }
	return d + "/.upharness/model-capabilities.json"
}

func loadCaps() map[string]int64 {
	out := map[string]int64{}
	data, err := os.ReadFile(capFile())
	if err != nil { return out }
	json.Unmarshal(data, &out)
	return out
}

func saveCaps(caps map[string]int64) {
	os.MkdirAll(os.Getenv("HOME")+"/.upharness", 0700)
	data, _ := json.Marshal(caps)
	os.WriteFile(capFile(), data, 0600)
}

func RecordFailure(id string) {
	caps := loadCaps()
	caps[id] = time.Now().Unix()
	saveCaps(caps)
}

func RecordSuccess(id string) {
	caps := loadCaps()
	delete(caps, id)
	saveCaps(caps)
}

func isUnhealthy(id string, caps map[string]int64) bool {
	ts, ok := caps[id]
	if !ok { return false }
	return time.Now().Unix()-ts < 6*3600
}

func LiveList() []ModelEntry {
	caps := loadCaps()
	seen := map[string]bool{}
	out := []ModelEntry{}

	if ids, err := NvidiaModels(); err == nil {
		for _, id := range ids {
			if !seen[id] {
				seen[id] = true
				out = append(out, ModelEntry{ID: id, Provider: "nvidia", Unhealthy: isUnhealthy(id, caps)})
			}
		}
	}
	if ids, err := OpenRouterModels(); err == nil {
		for _, id := range ids {
			if !seen[id] {
				seen[id] = true
				out = append(out, ModelEntry{ID: id, Provider: "openrouter", Unhealthy: isUnhealthy(id, caps)})
			}
		}
	}
	if len(out) == 0 { return StaticFallback() }
	sort.Slice(out, func(i, j int) bool {
		if out[i].Unhealthy != out[j].Unhealthy { return !out[i].Unhealthy }
		return out[i].ID < out[j].ID
	})
	return out
}

func TestModel(provider, id, key string) bool {
	var url, auth, body string
	if provider == "nvidia" {
		url = "https://integrate.api.nvidia.com/v1/chat/completions"
		auth = key
		body = "{"model":"" + id + "","messages":[{"role":"user","content":"OK"}],"max_tokens":1,"temperature":0}"
	} else if provider == "openrouter" {
		url = "https://openrouter.ai/api/v1/chat/completions"
		auth = os.Getenv("OPENROUTER_API_KEY")
		body = "{"model":"" + id + "","messages":[{"role":"user","content":"OK"}],"max_tokens":1,"temperature":0}"
	} else {
		return false
	}
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil { return false }
	req.Header.Set("Authorization", "Bearer "+auth)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil { return false }
	defer resp.Body.Close()
	if resp.StatusCode != 200 { return false }
	io.Copy(io.Discard, resp.Body)
	return true
}

func TestAll(provider string, ids []string, key string) []ModelEntry {
	caps := loadCaps()
	out := []ModelEntry{}
	for _, id := range ids {
		if isUnhealthy(id, caps) {
			out = append(out, ModelEntry{ID: id, Provider: provider, Unhealthy: true, LastCheck: caps[id]})
			continue
		}
		if TestModel(provider, id, key) {
			out = append(out, ModelEntry{ID: id, Provider: provider, Unhealthy: false, LastCheck: time.Now().Unix(), LatencyMs: 0})
			RecordSuccess(id)
		} else {
			out = append(out, ModelEntry{ID: id, Provider: provider, Unhealthy: true, LastCheck: time.Now().Unix()})
			RecordFailure(id)
		}
		time.Sleep(200 * time.Millisecond)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Unhealthy != out[j].Unhealthy { return !out[i].Unhealthy }
		return out[i].ID < out[j].ID
	})
	return out
}
