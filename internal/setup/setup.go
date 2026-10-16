package setup

import (
"encoding/json"
"fmt"
"io"
"net/http"
"os"
"path/filepath"
)

// Config holds everything the user pastes into the setup UI.
type Config struct {
	URL string
	Password string
	Telegram string
	NVIDIA string
	OpenRouter string
}

// Handler receives the pasted config, saves it locally and confirms.
// It never starts external binaries; services run in the same process
// or are started manually after configuration.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}
	var cfg Config
	if err := json.Unmarshal(body, &cfg); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	dir := filepath.Join(os.Getenv("HOME"), ".upharness")
	if dir == ".upharness" || dir == "/.upharness" {
		dir = ".upharness"
	}
	os.MkdirAll(dir, 0700)
	os.WriteFile(filepath.Join(dir, "config.json"), body, 0600)
	os.Setenv("UPHARNESS_URL", cfg.URL)
	os.Setenv("UPHARNESS_PASSWORD", cfg.Password)
	os.Setenv("TELEGRAM_BOT_TOKEN", cfg.Telegram)
	os.Setenv("NVIDIA_API_KEY", cfg.NVIDIA)
	os.Setenv("OPENROUTER_API_KEY", cfg.OpenRouter)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "{"status":"ok","message":"Configuration saved."}")
}
