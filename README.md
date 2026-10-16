# UPHarness

Autonomer AI-Harness mit Zero-Config-Setup. Einmal konfigurieren, alles startet automatisch.

UPHarness ist ein schlanker, plattformübergreifender Harness für autonome KI-Agenten. Du trägst einmal deine Zugangsdaten ein (Server-URL, Passwort, optionale API-Keys) – und alles andere startet von selbst.

## Wichtig: Fertige Binaries vs. Selbstbau

Aktuell gibt es noch keine fertigen EXE-Dateien zum direkten Herunterladen. Die GitHub Actions Builds laufen noch oder sind in der Warteschlange. Bis sie fertig sind, musst du UPHarness **selbst bauen**. Das ist unkompliziert und dauert wenige Minuten.

## Schritt 1: Go installieren

Lade Go von https://go.dev/dl/ herunter und installiere es für dein Betriebssystem.

- Windows: MSI-Installer ausführen
- macOS: PKG-Installer ausführen oder `brew install go`
- Linux: `sudo apt install golang-go` (oder je nach Distribution)
- Android (Termux): `pkg install golang git`

Überprüfe die Installation mit: `go version`

## Schritt 2: Repository herunterladen

Öffne ein Terminal (Eingabeaufforderung, PowerShell, Terminal oder Termux) und führe aus:

```bash
git clone https://github.com/HAX73w0rlD/upharness.git
cd upharness
```

## Schritt 3: Für dein Betriebssystem bauen

Führe im Ordner `upharness` den passenden Befehl aus:

| Betriebssystem | Befehl |
|----------------|--------|
| Windows | `GOOS=windows GOARCH=amd64 go build -o upharness.exe .` |
| Linux | `GOOS=linux GOARCH=amd64 go build -o upharness .` |
| macOS (Intel) | `GOOS=darwin GOARCH=amd64 go build -o upharness .` |
| macOS (Apple Silicon) | `GOOS=darwin GOARCH=arm64 go build -o upharness .` |
| Android (Termux) | `GOOS=android GOARCH=arm64 go build -o upharness .` |

Nach erfolgreichem Build liegt die Datei `upharness` (oder `upharness.exe`) im Ordner.

## Schritt 4: Starten

Starte die Binary im Terminal:

- Windows: `upharness.exe`
- Linux/macOS: `./upharness`
- Android (Termux): `./upharness`

Du siehst:

```
UPHarness Setup-Server läuft auf http://localhost:3000
Einfach URL & Passwort einfügen – alles startet automatisch
```

## Schritt 5: Einrichten im Browser

Öffne `http://localhost:3000` im Browser. Trage dort ein:

1. **OpenCode-Server URL** (z. B. `http://127.0.0.1:3002`)
2. **OpenCode Passwort**
3. Optional: **Telegram Bot Token**, **NVIDIA API Key**, **OpenRouter API Key**

Klicke auf „Speichern & automatisch starten“. Die Konfiguration wird lokal gespeichert und alle Dienste starten im Hintergrund.

## Android (Termux) – Komplettanleitung

1. Termux aus dem **F-Droid Store** installieren (nicht Play Store)
2. Termux öffnen und aktualisieren: `pkg update && pkg upgrade`
3. Go und Git installieren: `pkg install golang git`
4. Repository klonen: `git clone https://github.com/HAX73w0rlD/upharness.git && cd upharness`
5. Bauen: `GOOS=android GOARCH=arm64 go build -o upharness .`
6. Starten: `./upharness`
7. Browser öffnen: `http://localhost:3000`

Wichtig: Android-Batterieoptimierung für Termux ausschalten, sonst stirbt der Prozess im Hintergrund.

## Später: Fertige Binaries von GitHub Actions

Sobald die Builds auf https://github.com/HAX73w0rlD/upharness/actions erfolgreich durchgelaufen sind, kannst du dort fertige Binaries für Windows, Linux, macOS und Android herunterladen und den Selbstbau überspringen.

## Sicherheit

Keine API-Keys oder Passwörter werden ins Repository gepusht. Konfiguration wird lokal unter `~/.upharness/config.json` gespeichert.

Als publizierte Modell-Anbieter sind **NVIDIA** und **OpenRouter** hinterlegt – ohne API-Werte im Code.

## Lizenz

MIT

---

UPHarness, entwickelt von Fudasys.