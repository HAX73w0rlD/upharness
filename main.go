package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HAX73w0rlD/upharness/internal/setup"
)

func main() {
	http.HandleFunc("/", setupUIHandler)
	http.HandleFunc("/api/setup", setup.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Ì†ºÌæ® UPHarness Setup-Server l√§uft auf http://localhost:" + port)
	fmt.Println("Ì†ΩÌ≥ù Einfach URL & Passwort einf√ºgen ‚Äì alles startet automatisch")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupUIHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/setup.html")
}
