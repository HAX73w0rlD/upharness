# UPHarness

Autonomer AI-Harness mit Zero-Config-Setup. Einmal konfigurieren, alles startet automatisch.

UPHarness ist ein schlanker, plattformübergreifender Harness für autonome KI-Agenten. Du trägst einmal deine Zugangsdaten ein (Server-URL, Passwort, optionale API-Keys) – und alles andere startet von selbst. Kein manuelles Konfigurieren von Umgebungsvariablen, keine Konfigurationsdateien suchen, keine Dutzende Seiten Dokumentation lesen vor dem ersten Start.

## Was du brauchst

- Ein Gerät mit Windows, Linux, macOS oder Android (Termux)
- Einen OpenCode-Server, den du bereits betreibst (oder den du aufsetzt)

## Schnellstart (alle Plattformen)

### Schritt 1: Herunterladen oder bauen

Du hast zwei Möglichkeiten, UPHarness auf dein Gerät zu bekommen:

**A) Fertige Binary von GitHub Actions (einfachste Variante)**
1. Gehe auf https://github.com/HAX73w0rlD/upharness/actions
2. Klicke auf den neuesten erfolgreichen Build für dein Betriebssystem
3. Lade die Artifact-Datei herunter (z. B. `upharness-windows-amd64.exe`, `upharness-linux-amd64`, `upharness-macos-arm64`)
4. Entpacke sie an einem Ort deiner Wahl

**B) Selbst bauen (falls du Go installiert hast)**
1. Installiere Go ab https://go.dev/dl/
2. Öffne ein Terminal und klone das Repository:
   ```bash
   git clone https://github.com/HAX73w0rlD/upharness.git
   cd upharness
   ```
3. Baue für dein System:
   - Windows: `GOOS=windows GOARCH=amd64 go build -o upharness.exe .`
   - Linux: `GOOS=linux GOARCH=amd64 go build -o upharness .`
   - macOS: `GOOS=darwin GOARCH=arm64 go build -o upharness .`
   - Android (Termux): `GOOS=android GOARCH=arm64 go build -o upharness .`


### Schritt 2: Starten

Öffne ein Terminal (oder die Eingabeaufforderung / PowerShell unter Windows) im Ordner mit der `upharness`-Datei und starte sie:

- Windows: `upharness.exe`
- Linux/macOS: `./upharness`
- Android (Termux): `./upharness`

Du siehst eine Meldung wie:

```
UPHarness Setup-Server läuft auf http://localhost:3000
Einfach URL & Passwort einfügen – alles startet automatisch
```

### Schritt 3: Einrichten im Browser

Öffne im Browser `http://localhost:3000`. Trage dort ein:

1. **OpenCode-Server URL** (z. B. `http://127.0.0.1:3002`)
2. **OpenCode Passwort**
3. Optional: **Telegram Bot Token**, **NVIDIA API Key**, **OpenRouter API Key**

Klicke auf „Speichern & automatisch starten“. UPHarness speichert die Konfiguration lokal und startet im Hintergrund Web-Server, Daemon und Telegram-Bridge.


## Android (Termux) – Schritt für Schritt

1. **Termux installieren**: Lade Termux aus dem F-Droid Store herunter (nicht aus dem Play Store, der ist veraltet).
2. **Termux öffnen** und Pakete aktualisieren:
   ```bash
   pkg update && pkg upgrade
   ```
3. **Go installieren**:
   ```bash
   pkg install golang git
   ```
4. **Repository klonen**:
   ```bash
   git clone https://github.com/HAX73w0rlD/upharness.git
   cd upharness
   ```
5. **Für Android bauen**:
   ```bash
   GOOS=android GOARCH=arm64 go build -o upharness .
   ```
6. **Starten**:
   ```bash
   ./upharness
   ```
7. **Im Browser öffnen**: Gehe in einem Android-Browser zu `http://localhost:3000` und richte UPHarness ein wie oben beschrieben.

Wichtig: Stelle sicher, dass Termux im Hintergrund laufen darf (Android-Batterieoptimierung für Termux ausschalten), sonst stirbt der Prozess, wenn der Bildschirm ausgeht.


## Was passiert nach dem Start?

Nachdem du im Browser auf „Speichern & automatisch starten“ geklickt hast, passiert Folgendes automatisch:

- Deine Zugangsdaten werden lokal gespeichert (unter `~/.upharness/config.json`, **nie** ins Repository gepusht)
- Die Environment-Variablen werden gesetzt
- Der Web-Server, der Daemon und optional die Telegram-Bridge starten im Hintergrund
- Du kannst das Terminal-Fenster schließen, UPHarness läuft weiter


## Sicherheit

Keine API-Keys oder Passwörter werden ins Repository gepusht. Die `.gitignore` schließt alle Secret-Dateien aus. Konfiguration wird ausschließlich lokal auf deinem Gerät gespeichert und über Environment-Variablen an die Dienste weitergereicht.

Als publizierte Modell-Anbieter sind derzeit **NVIDIA** und **OpenRouter** hinterlegt – ohne API-Werte im Code.


## Lizenz

MIT

---

UPHarness, entwickelt von Fudasys.