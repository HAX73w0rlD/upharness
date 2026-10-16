package main

import (
"fmt"
"log"
"net/http"
"os"

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
	case "session":
		fmt.Println("Session management - not yet implemented")
	case "models":
		fmt.Println("Available models: nvidia, opencode (free tiers)")
	case "telegram":
		fmt.Println("Telegram bridge - configure token via web setup first")
	case "daemon":
		fmt.Println("Background daemon - not yet implemented")
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Usage: upharness [web|session|models|telegram|daemon]")
		fmt.Println("No arguments starts the setup web server.")
		os.Exit(1)
	}
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
