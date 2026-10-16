# Ì†ºÌæ® UPHarness ‚Äì Lightweight AI-Harness

**Einfach. Sch√∂n. Automatisch.**

UPHarness ist ein schlanker, plattform√ºbergreifender AI-Harness. Du f√ºgst deine Links und Keys in die Web-Oberfl√§che ein ‚Äì alles andere startet automatisch.

## ‚ú® Features

| Feature | Beschreibung |
|---------|-------------|
| Ì†ºÌæ® **Sch√∂nes Setup** | Glas-Design, Gradienten, einfach zu bedienen |
| Ì†ΩÌ¥ó **Einf√ºgen & Starten** | Links/Token einf√ºgen, alles startet automatisch |
| Ì†ΩÌ∂•Ô∏è **Cross-Platform** | Windows, Linux, macOS, Android |
| ‚ö° **Leichtgewichtig** | In Go, nativ, geringer Speicherverbrauch |
| Ì†æÌ¥ñ **Telegram-Ready** | Optional: Bot Token einf√ºgen |
| Ì†ºÌ∂ì **Kostenlose Modelle** | NVIDIA & OpenRouter unterst√ºtzen kostenlose Tiers |

## Ì†ΩÌ∫Ä Schnellstart

### 1. Build & Start
```bash
go build -o upharness .
./upharness
```

### 2. Browser √∂ffnen
√ñffne `http://localhost:3000`

### 3. Einf√ºgen & Starten
F√ºge in die Web-Oberfl√§che ein:
- OpenCode-Server URL (z. B. `http://127.0.0.1:3002`)
- OpenCode Passwort
- Optional: Telegram Bot Token, NVIDIA Key, OpenRouter Key

Klicke auf **‚ÄûAlles speichern & automatisch starten‚Äú** ‚Äì fertig.

## Ì†ΩÌ≥± Plattformen

| Plattform | Build |
|-----------|-------|
| Windows | `GOOS=windows GOARCH=amd64 go build` |
| Linux | `GOOS=linux GOARCH=amd64 go build` |
| macOS | `GOOS=darwin GOARCH=arm64 go build` |
| Android | `GOOS=android GOARCH=arm64 go build` |

## Ì†ΩÌ¥ê Sicherheit

Alle Keys werden lokal unter `~/.upharness/config.json` gespeichert und **nie** ins Git-Repository gepusht (siehe `.gitignore`).

## Ì†ΩÌ≥Ñ Lizenz

MIT
