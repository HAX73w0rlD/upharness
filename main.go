package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("UPHarness - lightweight AI harness")
		fmt.Println("Usage: upharness <command> [args]")
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  session   Manage sessions")
		fmt.Println("  models    List available models")
		fmt.Println("  telegram  Start Telegram bridge")
		fmt.Println("  web       Start web server")
		fmt.Println("  daemon    Start background daemon")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "session":
		fmt.Println("Session management")
	case "models":
		fmt.Println("Available models")
	case "telegram":
		fmt.Println("Telegram bridge")
	case "web":
		fmt.Println("Web server")
	case "daemon":
		fmt.Println("Background daemon")
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}
