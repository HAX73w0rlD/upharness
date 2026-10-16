package setup

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// Config is everything the user pastes into the setup UI.
type Config struct {
	URL        string
	Password   string
	Telegram   string
	NVIDIA     string
	OpenRouter string
}

// Handler receives the pasted config and starts everything automatically.
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

	// 1. Save config to .upharness (local only, never pushed)
	dir := filepath.Join(os.Getenv("HOME"), ".upharness")
	os.MkdirAll(dir, 0700)
	os.WriteFile(filepath.Join(dir, "config.json"), body, 0600)

	// 2. Export env keys (only used keys)
	os.Setenv("UPHARNESS_URL", cfg.URL)
	os.Setenv("UPHARNESS_PASSWORD", cfg.Password)
	os.Setenv("TELEGRAM_BOT_TOKEN", cfg.Telegram)
	os.Setenv("NVIDIA_API_KEY", cfg.NVIDIA)
	os.Setenv("OPENROUTER_API_KEY", cfg.OpenRouter)

	// 3. Auto-start: launch daemon + web + telegram in background
	go startServices()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status":"ok","message":"Configuration saved. UPHarness starting automatically."}`)
}

func startServices() {
	cmds := [][]string{
		{"go", "run", "./cmd/web"},
		{"go", "run", "./cmd/daemon"},
		{"go", "run", "./cmd/telegram"},
	}
	for _, c := range cmds {
		cmd := exec.Command(c[0], c[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
	}
}
