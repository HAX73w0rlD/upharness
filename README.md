# UPHarness

Autonomer AI-Harness mit Zero-Config-Setup. Einmal konfigurieren, alles startet automatisch.

UPHarness ist ein schlanker, plattformuebergreifender Harness fuer autonome KI-Agenten. Du traegst einmal deine Zugangsdaten ein (Server-URL, Passwort, optionale API-Keys) und alles andere startet von selbst.

## Stand des Projekts

Dies ist eine Beta. Es gibt noch keine fertigen Binaries zum Herunterladen, die GitHub Actions Builds sind noch nicht stabil. Der funktionierende Weg ist der Selbstbau mit Go (dauert wenige Minuten). Sobald die Actions gruene Builds liefern, kannst du fertige Binaries unter https://github.com/HAX73w0rlD/upharness/actions herunterladen.

## Voraussetzungen

- Go (https://go.dev/dl/), Version 1.23 oder neuer
- Git
- Einen OpenCode-Server, den du bereits betreibst

Pruefen mit: go version und git --version

## Installation Schritt fuer Schritt

### 1. Repository holen

```
git clone https://github.com/HAX73w0rlD/upharness.git
cd upharness
```

### 2. Bauen

```
GOOS=linux GOARCH=amd64 go build -o upharness .     (Linux)
GOOS=windows GOARCH=amd64 go build -o upharness.exe . (Windows)
GOOS=darwin GOARCH=arm64 go build -o upharness .    (macOS Apple Silicon)
GOOS=darwin GOARCH=amd64 go build -o upharness .    (macOS Intel)
GOOS=android GOARCH=arm64 go build -o upharness .   (Android/Termux)
```

### 3. Starten

```
./upharness            (Linux/macOS/Android)
upharness.exe          (Windows)
```

Ohne Argumente startet der Setup-Webserver auf Port 3000. Alternativ: ./upharness web.

### 4. Einrichten im Browser

Oeffne http://localhost:3000 und trage ein:

- OpenCode-Server URL (z. B. http://127.0.0.1:3002)
- OpenCode Passwort
- Optional: Telegram Bot Token, NVIDIA API Key, OpenRouter API Key

Klicke auf Speichern. Die Konfiguration landet nur lokal unter ~/.upharness/config.json.

### 5. Modelle pruefen

```
./upharness models
```

Ohne Keys siehst du nur die Fallback-Liste. Mit NVIDIA_API_KEY gesetzt wird die Live-Liste von https://integrate.api.nvidia.com/v1/models geholt, dazu die oeffentliche OpenRouter-Liste von https://openrouter.ai/api/v1/models. Modelle, die kuerzlich fehlgeschlagen sind, werden 6 Stunden als [unhealthy, cached] markiert statt sie zu verstecken.

```
export NVIDIA_API_KEY=nvapi-dein-key
./upharness models
```

Hinweis: Ein gelistetes Modell ist keine Garantie, dass es gerade Kapazitaet hat. Fehler werden in ~/.upharness/model-capabilities.json zwischengespeichert.

### 6. Aktualisieren (Auto-Update)

```
./upharness update       (holt git pull und zeigt den passenden Build-Befehl)
sh scripts/update.sh     (pull + Build in einem Schritt)
```

Danach die neue Binary neu starten.

## Android (Termux) Komplettanleitung

1. Termux aus dem F-Droid Store installieren (nicht Play Store).
2. Termux oeffnen: pkg update && pkg upgrade
3. Go und Git installieren: pkg install golang git
4. Klonen: git clone https://github.com/HAX73w0rlD/upharness.git && cd upharness
5. Bauen: GOOS=android GOARCH=arm64 go build -o upharness .
6. Starten: ./upharness
7. Im Handy-Browser http://localhost:3000 oeffnen.

Wichtig: Android-Batterieoptimierung fuer Termux ausschalten, sonst stirbt der Prozess im Hintergrund.

## Fehlerbehebung

- invalid UTF-8 encoding beim Build: alte main.go ausgecheckt. Loesung: git pull, dann neu bauen. Die aktuelle main.go ist reines ASCII.
- Unknown command: /pfad/upharness: alte Binary ohne Argument-Unterstuetzung. Loesung: git pull + neu bauen. Aktuell startet ./upharness ohne Argumente den Server.
- Port 3000 belegt: eine alte Instanz laeuft noch. Loesung: pkill -f upharness oder PORT=8080 ./upharness und dann http://localhost:8080 oeffnen.
- Keine Modelle / nur Fallback: NVIDIA_API_KEY nicht gesetzt oder kein Netz. Loesung: export NVIDIA_API_KEY=... und ./upharness models erneut.
- GLM antwortet nicht: Provider meldet temporaere Kapazitaetsprobleme. Die Liste zeigt es weiter an, markiert es aber nach Fehlschlag als unhealthy. Alternative aus der Liste waehlen.
- Actions ohne Artifacts: Builds haengen in queued/running. Dann Selbstbau verwenden, bis ein gruener Build mit Artifacts vorliegt.

## Sicherheit

Keine API-Keys oder Passwoerter werden ins Repository gepusht. Konfiguration liegt nur lokal unter ~/.upharness/. Publiziert sind nur die Provider-Namen nvidia und opencode/openrouter, nie Schluesselwerte.

## Lizenz

MIT

---

UPHarness, entwickelt von Fudasys.