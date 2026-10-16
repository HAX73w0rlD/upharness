package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/HAX73w0rlD/upharness/internal/catalog"
	"github.com/HAX73w0rlD/upharness/internal/setup"
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
		fmt.Println("Session management - kommende Version")
	case "telegram":
		fmt.Println("Telegram bridge - configure token via web setup first")
	case "daemon":
		fmt.Println("Background daemon - kommende Version")
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
		if m.Unhealthy {
			mark = " [nicht antwortend, 6h Pause]"
		}
		fmt.Printf(" - %s (%s)%s
", m.ID, m.Provider, mark)
	}
	if os.Getenv("NVIDIA_API_KEY") == "" {
		fmt.Println("
Hinweis: NVIDIA_API_KEY nicht gesetzt - NVIDIA Live-Liste übersprungen.")
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
	fmt.Println("
Repository aktualisiert. Jetzt neu bauen:")
	fmt.Printf("  GOOS=%s GOARCH=%s go build -o upharness .
", runtime.GOOS, runtime.GOARCH)
	fmt.Println("  Danach: ./upharness (oder upharness.exe unter Windows)")
}

func runDoctor() {
	fmt.Println("=== UPHarness Diagnose ===")
	fmt.Printf("OS: %s %s (%s)
", runtime.GOOS, runtime.GOARCH, runtime.Version())
	fmt.Printf("NVIDIA_API_KEY: %s
", map[bool]string{true: "gesetzt", false: "nicht gesetzt"}[os.Getenv("NVIDIA_API_KEY") != ""])
	fmt.Printf("OPENROUTER_API_KEY: %s
", map[bool]string{true: "gesetzt", false: "nicht gesetzt"}[os.Getenv("OPENROUTER_API_KEY") != ""])
	fmt.Println("
Modell-Check:")
	printModels()
}

func startServer() {
	http.HandleFunc("/", setupUIHandler)
	http.HandleFunc("/api/setup", setup.Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("UPHarness Setup-Server läuft auf http://localhost:" + port)
	fmt.Println("Öffne diese Adresse im Browser zum Konfigurieren.")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupUIHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/setup.html")
}
