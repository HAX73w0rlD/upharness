# UPHarness

Autonomer AI-Harness mit Zero-Config-Setup. Einmal konfigurieren, alles startet automatisch.

UPHarness ist ein schlanker, plattformuebergreifender Harness für autonome KI-Agenten. Du trägst einmal deine Zugangsdaten ein (Server-URL, Passwort, optionale API-Keys) – und alles andere startet von selbst.

## Plattformen

| Plattform | Build |
|-----------|-------|
| Windows | GOOS=windows GOARCH=amd64 go build -o upharness.exe . |
| Linux | GOOS=linux GOARCH=amd64 go build -o upharness . |
| macOS (Apple Silicon) | GOOS=darwin GOARCH=arm64 go build -o upharness . |
| macOS (Intel) | GOOS=darwin GOARCH=amd64 go build -o upharness . |
| Android (Termux) | GOOS=android GOARCH=arm64 go build -o upharness . |

## Schnellstart

### 1. Go installieren
Von https://go.dev/dl/ oder unter Termux: pkg install golang git

### 2. Repository klonen und bauen
git clone https://github.com/HAX73w0rlD/upharness.git
cd upharness
GOOS=linux GOARCH=amd64 go build -o upharness .
./upharness

### 3. Einrichten
Öffne http://localhost:3000 im Browser. Trage ein:
- OpenCode-Server URL (z. B. http://127.0.0.1:3002)
- OpenCode Passwort
- Optional: Telegram Bot Token, NVIDIA API Key, OpenRouter API Key

Klicke auf Speichern & automatisch starten.

### 4. Modelle prüfen
./upharness models

Zeigt nur Modelle, die aktuell antworten. Nicht-antwortende werden 6h pausiert.

### 5. Aktualisieren
./upharness update
sh scripts/update.sh

## Sicherheit

Keine API-Keys oder Passwörter werden ins Repository gepusht. Konfiguration wird lokal unter ~/.upharness/config.json gespeichert.

## Lizenz

MIT

---

UPHarness, entwickelt von Fudasys.
