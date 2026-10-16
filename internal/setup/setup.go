package setup

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	URL       string
	Password  string
	Telegram  string
	NVIDIA    string
	OpenRouter string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST erforderlich", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil { http.Error(w, "schlechter Body", http.StatusBadRequest); return }
	var cfg Config
	if err := json.Unmarshal(body, &cfg); err != nil { http.Error(w, "schlechtes JSON", http.StatusBadRequest); return }

	home := os.Getenv("HOME")
	if home == "" { home = "." }
	dir := filepath.Join(home, ".upharness")
	os.MkdirAll(dir, 0700)
	os.WriteFile(filepath.Join(dir, "config.json"), body, 0600)

	os.Setenv("UPHARNESS_URL", cfg.URL)
	os.Setenv("UPHARNESS_PASSWORD", cfg.Password)
	os.Setenv("TELEGRAM_BOT_TOKEN", cfg.Telegram)
	os.Setenv("NVIDIA_API_KEY", cfg.NVIDIA)
	os.Setenv("OPENROUTER_API_KEY", cfg.OpenRouter)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok: Konfiguration gespeichert. UPHarness startet automatisch.")
}
