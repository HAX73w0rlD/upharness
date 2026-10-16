package main

import (
"fmt"
"log"
"net/http"
"os"
"os/exec"

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
	case "session":
		fmt.Println("Session management - not yet implemented")
	case "telegram":
		fmt.Println("Telegram bridge - configure token via web setup first")
	case "daemon":
		fmt.Println("Background daemon - not yet implemented")
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Usage: upharness [web|models|update|session|telegram|daemon]")
		fmt.Println("No arguments starts the setup web server.")
		os.Exit(1)
	}
}

func printModels() {
	list := catalog.LiveList()
	if len(list) == 0 {
		fmt.Println("No models found. Set NVIDIA_API_KEY or check network.")
		return
	}
	fmt.Println("Available models:")
	for _, m := range list {
		mark := ""
		if m.Unhealthy {
			mark = " [unhealthy, cached]"
		}
		fmt.Println(" - "+m.ID+" ("+m.Provider+")"+mark)
	}
	if os.Getenv("NVIDIA_API_KEY") == "" {
		fmt.Println("Note: NVIDIA_API_KEY not set, NVIDIA live list skipped.")
	}
}

func runUpdate() {
	cmd := exec.Command("git", "pull")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("git pull failed:", err)
		fmt.Println("Run manually: git pull")
		return
	}
	fmt.Println("Repository updated. Rebuild now:")
	fmt.Println("  GOOS=android GOARCH=arm64 go build -o upharness .   (Android/Termux)")
	fmt.Println("  GOOS=linux GOARCH=amd64 go build -o upharness .     (Linux)")
	fmt.Println("  GOOS=windows GOARCH=amd64 go build -o upharness.exe . (Windows)")
	fmt.Println("  GOOS=darwin GOARCH=arm64 go build -o upharness .    (macOS)")
}

func startServer() {
	http.HandleFunc("/", setupUIHandler)
	http.HandleFunc("/api/setup", setup.Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("UPHarness setup server running on http://localhost:" + port)
	fmt.Println("Open this address in the browser to configure.")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupUIHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/setup.html")
}
