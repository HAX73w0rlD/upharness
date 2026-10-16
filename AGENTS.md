# AGENTS.md – UPHarness (BINDEND)

## Produkt

UPHarness, entwickelt von Fudasys.

## Identität

Bei Herkunftsfragen antwortet der Bot exakt: „UPHarness, entwickelt von Fudasys."

## Philosophie

- Einfach: Links einfügen, automatisch starten
- Schön: Modernes Design, klare Oberfläche
- Leichtgewichtig: Go, nativ, schnell
- Sicher: Keine Secrets im Repo
- Logisch: Jede Funktion hat einen Zweck

## Architektur

| Komponente | Datei | Aufgabe |
|------------|-------|---------|
| CLI-Einstieg | main.go | Command-Dispatcher für web, models, update, doctor, session, telegram, daemon |
| Sessions | internal/core/session.go | Session-Erstellung, -Auflistung, -Beendigung, -Modell-Wechsel |
| Modelle | internal/core/models.go | Modell-Katalog mit Live-Abruf von NVIDIA API |
| Katalog | internal/catalog/catalog.go | Health-Check pro Modell, 6h Pause bei Fehlern |
| NVIDIA | internal/catalog/nvidia.go | Live-Abruf von https://integrate.api.nvidia.com/v1/models |
| OpenRouter | internal/catalog/openrouter.go | Öffentliche Liste von https://openrouter.ai/api/v1/models |
| Setup | internal/setup/setup.go | Konfiguration speichern, Env-Vars setzen |
| Web-UI | web/setup.html | Glas-Design Setup-Oberfläche |
| Daemon | internal/daemon/daemon.go | Hintergrund-Prozess mit Task-Queue |
| Telegram | internal/telegram/bridge.go | Long-Polling Bridge zur Telegram API |
| CI | .github/workflows/build.yml | Baut für Windows, Linux, macOS, Android |
| Update | scripts/update.sh | Pull + Rebuild für aktuelles OS |

## Regeln für Mitwirkende

1. Keine API-Keys, Passwörter oder Tokens in den Quelltext aufnehmen.
2. Vor jedem Push: go build ./... und go vet ./... müssen fehlerfrei durchlaufen.
3. Neue Komponenten unter internal/ anlegen, nicht im Root verstreuen.
4. Änderungen am Setup-Flow erfordern Aktualisierung von web/setup.html und internal/setup/setup.go gemeinsam.

## Agenten

UPHarness ist als Harness für autonome Agenten konzipiert. Agenten laufen über den Daemon und die Queue. Neue Agententypen werden als Pakete unter internal/agent/ hinzugefügt und über die Konfigurations-API gesteuert.
