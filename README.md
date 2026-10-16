# UPHarness

Autonomer AI-Harness mit Zero-Config-Setup.

UPHarness ist ein schlanker, plattformübergreifender Harness für autonome KI-Agenten. Die Idee ist einfach: Du trägst einmal deine Zugangsdaten ein – Server-URL, Passwort, optionale API-Keys – und alles andere startet von selbst.

## Für wen ist das gedacht

UPHarness richtet sich an Entwickler und Teams, die KI-Agenten lokal oder auf eigenen Servern betreiben wollen, ohne sich mit Infrastruktur-Konfiguration aufzuhalten. Es läuft auf Windows, Linux, macOS und Android.

## Schnellstart

### 1. Bauen

```bash
go build -o upharness .
```

### 2. Starten

```bash
./upharness
```

### 3. Einrichten

Öffne im Browser http://localhost:3000. Trage dort ein:

- OpenCode-Server URL (z. B. http://127.0.0.1:3002)
- OpenCode Passwort
- Optional: Telegram Bot Token, NVIDIA API Key, OpenRouter API Key

Klicke auf „Speichern & automatisch starten“. UPHarness schreibt die Konfiguration lokal unter ~/.upharness/config.json und startet im Hintergrund Web-Server, Daemon und Telegram-Bridge.

## Architektur

UPHarness besteht aus lose gekoppelten Komponenten:

| Komponente | Aufgabe |
|------------|---------|
| Web-Server | Bereitstellt das Setup-Frontend und die Konfigurations-API |
| Setup | Nimmt Benutzereingaben entgegen und startet alle Dienste automatisch |
| Core | Session-Management und Modell-Katalog |
| Queue | Hintergrund-Aufgaben mit Persistenz |
| Daemon | Langlaufender Worker für wiederkehrende Aufgaben |
| Telegram | Optional: Anbindung an Telegram-Bots |
| Identity | Produkterkennung und Herkunftsanfragen |

Alle Komponenten sind in Go geschrieben und bauen zu einer einzigen nativen Binary.

## Plattformen

Native Builds über GitHub Actions:

- Windows (GOOS=windows GOARCH=amd64)
- Linux (GOOS=linux GOARCH=amd64)
- macOS (GOOS=darwin GOARCH=arm64)
- Android (GOOS=android GOARCH=arm64)

## Sicherheit

Keine API-Keys oder Passwörter werden ins Repository gepusht. Die .gitignore schließt alle Secret-Dateien aus (.upharness/, secrets.env, service.json, Logs, PIDs). Konfiguration wird ausschließlich lokal gespeichert und über Environment-Variablen an die Dienste weitergereicht.

Als publizierte Modell-Anbieter sind derzeit NVIDIA und OpenRouter hinterlegt – ohne API-Werte im Code.

## Lizenz

MIT

---

UPHarness, entwickelt von Fudasys.