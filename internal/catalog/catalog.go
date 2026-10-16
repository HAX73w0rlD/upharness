package catalog

import (
"encoding/json"
"os"
"path/filepath"
"sort"
"time"
)

// Provider lists only publishable providers. No keys here.
type Provider struct {
	ID string
	Free bool
}

func Providers() []Provider {
	return []Provider{{ID: "nvidia", Free: true}, {ID: "opencode", Free: true}}
}

// ModelEntry is one model with health state.
type ModelEntry struct {
	ID string
	Provider string
	Unhealthy bool
}

// StaticFallback is used when no network or no keys are available.
func StaticFallback() []ModelEntry {
	return []ModelEntry{
		{ID: "z-ai/glm-4.6", Provider: "nvidia"},
		{ID: "meta/llama-3.1-8b-instruct", Provider: "openrouter"},
	}
}

// capFile stores last-failure timestamps per model ID.
func capFile() string {
	d := os.Getenv("HOME")
	if d == "" {
		d = "."
	}
	return filepath.Join(d, ".upharness", "model-capabilities.json")
}

func loadCaps() map[string]int64 {
	out := map[string]int64{}
	data, err := os.ReadFile(capFile())
	if err != nil {
		return out
	}
	json.Unmarshal(data, &out)
	return out
}

func saveCaps(caps map[string]int64) {
	os.MkdirAll(filepath.Dir(capFile()), 0700)
	data, err := json.Marshal(caps)
	if err != nil {
		return
	}
	os.WriteFile(capFile(), data, 0600)
}

// RecordFailure marks a model as recently failed (unhealthy for 6h).
func RecordFailure(id string) {
	caps := loadCaps()
	caps[id] = time.Now().Unix()
	saveCaps(caps)
}

// RecordSuccess clears the failure mark of a model.
func RecordSuccess(id string) {
	caps := loadCaps()
	delete(caps, id)
	saveCaps(caps)
}

func isUnhealthy(id string, caps map[string]int64) bool {
	ts, ok := caps[id]
	if !ok {
		return false
	}
	return time.Now().Unix()-ts < 6*3600
}

// LiveList merges live NVIDIA + OpenRouter catalogs.
// Without keys or network it falls back to the static list.
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
	if len(out) == 0 {
		return StaticFallback()
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Unhealthy != out[j].Unhealthy {
			return !out[i].Unhealthy
		}
		return out[i].ID < out[j].ID
	})
	return out
}
