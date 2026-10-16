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

## Regeln für Mitwirkende

1. Keine API-Keys, Passwörter oder Tokens in den Quelltext aufnehmen.
2. Vor jedem Push: go build ./... und go vet ./... müssen fehlerfrei durchlaufen.
3. Neue Komponenten unter internal/ anlegen, nicht im Root verstreuen.
4. Änderungen am Setup-Flow erfordern Aktualisierung von web/setup.html und internal/setup/setup.go gemeinsam.

## Agenten

UPHarness ist als Harness für autonome Agenten konzipiert. Agenten laufen über den Daemon und die Queue. Neue Agententypen werden als Pakete unter internal/agent/ hinzugefügt und über die Konfigurations-API gesteuert.
