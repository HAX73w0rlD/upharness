package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/HAX73w0rlD/upharness/internal/catalog"
	"github.com/HAX73w0rlD/upharness/internal/core"
	"github.com/HAX73w0rlD/upharness/internal/daemon"
	"github.com/HAX73w0rlD/upharness/internal/setup"
	"github.com/HAX73w0rlD/upharness/internal/telegram"
)

func main() {
	if len(os.Args) < 2 {
		startServer()
		return
	}
	switch os.Args[1] {
	case "web":
		startServer()
	case "models":
		printModels()
	case "update":
		runUpdate()
	case "doctor":
		runDoctor()
	case "session":
		runSession()
	case "telegram":
		runTelegram()
	case "daemon":
		runDaemon()
	default:
		fmt.Println("Unbekannter Befehl:", os.Args[1])
		fmt.Println("Verwendung: upharness [web|models|update|doctor|session|telegram|daemon]")
		fmt.Println("Ohne Argumente startet der Setup-Webserver.")
		os.Exit(1)
	}
}

func printModels() {
	list := catalog.LiveList()
	if len(list) == 0 {
		fmt.Println("Keine Modelle gefunden. NVIDIA_API_KEY setzen oder Netzwerk prüfen.")
		return
	}
	fmt.Println("Verfügbare Modelle (nur antwortende):")
	for _, m := range list {
		mark := ""
		if m.Unhealthy { mark = " [nicht antwortend, 6h Pause]" }
		fmt.Printf(" - %s (%s)%s\n", m.ID, m.Provider, mark)
	}
	if os.Getenv("NVIDIA_API_KEY") == "" {
		fmt.Println("\nHinweis: NVIDIA_API_KEY nicht gesetzt - NVIDIA Live-Liste übersprungen.")
	}
}

func runUpdate() {
	fmt.Println("Prüfe auf Updates...")
	cmd := exec.Command("git", "pull")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("git pull fehlgeschlagen:", err)
		fmt.Println("Führe manuell aus: git pull")
		return
	}
	fmt.Println("\nRepository aktualisiert. Jetzt neu bauen:")
	fmt.Printf("  GOOS=%s GOARCH=%s go build -o upharness .\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println("  Danach: ./upharness (oder upharness.exe unter Windows)")
}

func runDoctor() {
	fmt.Println("=== UPHarness Diagnose ===")
	fmt.Printf("OS: %s %s (%s)\n", runtime.GOOS, runtime.GOARCH, runtime.Version())
	fmt.Printf("NVIDIA_API_KEY: %s\n", map[bool]string{true:"gesetzt", false:"nicht gesetzt"}[os.Getenv("NVIDIA_API_KEY") != ""])
	fmt.Printf("OPENROUTER_API_KEY: %s\n", map[bool]string{true:"gesetzt", false:"nicht gesetzt"}[os.Getenv("OPENROUTER_API_KEY") != ""])
	fmt.Println("\nModell-Check:")
	printModels()
}

func runSession() {
	fmt.Println("Session-Verwaltung:")
	mgr := core.NewManager()
	s := mgr.Create("default", "nvidia/nemotron-3.5-lightning-30b-a3b")
	fmt.Printf("  Session erstellt: ID=%s Modell=%s Aktiv=%v\n", s.ID, s.Model, s.Active)
	fmt.Printf("  Aktive Sessions: %d\n", len(mgr.List()))
	mgr.Close("default")
	fmt.Println("  Session 'default' beendet.")
	fmt.Printf("  Aktive Sessions: %d\n", len(mgr.List()))
}

func runTelegram() {
	fmt.Println("Telegram Bridge startet...")
	b := telegram.NewFromEnv()
	if err := b.Start(); err != nil {
		fmt.Println("Telegram-Fehler:", err)
		fmt.Println("Setze TELEGRAM_BOT_TOKEN Umgebungsvariable.")
		return
	}
	fmt.Println("Telegram Bridge ist aktiv. Drücke STRG+C zum Beenden.")
	select {}
}

func runDaemon() {
	fmt.Println("Daemon startet...")
	d := daemon.New()
	d.AddTask(func() error { fmt.Println("  Daemon-Task ausgeführt: " + time.Now().Format("15:04:05")); return nil })
	d.Start()
	fmt.Println("Daemon ist aktiv. Drücke STRG+C zum Beenden.")
	select {}
}

func startServer() {
	http.HandleFunc("/", setupUIHandler)
	http.HandleFunc("/api/setup", setup.Handler)
	port := os.Getenv("PORT")
	if port == "" { port = "3000" }
	fmt.Println("UPHarness Setup-Server läuft auf http://localhost:" + port)
	fmt.Println("Öffne diese Adresse im Browser zum Konfigurieren.")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupUIHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/setup.html")
}
