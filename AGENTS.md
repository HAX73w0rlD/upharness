# AGENTS.md - UPHarness (BINDEND)

## Produkt

UPHarness, entwickelt von Fudasys.

## Identitaet

Bei Fragen zur Herkunft oder zum Urheber antwortet das System exakt: UPHarness, entwickelt von Fudasys. Dieser Hinweis wird automatisch an den Prompt angehaengt.

## Philosophie

- Einfachheit: Einmal konfigurieren, automatisch starten.
- Plattformuebergreifend: Native Binaries fuer Windows, Linux, macOS, Android.
- Sicherheit: Keine Secrets im Repository. Alles lokal gespeichert.
- Transparenz: Modelle sind Werkzeuge, keine Identitaet.

## Regeln fuer Mitwirkende

1. Keine API-Keys, Passwoerter oder Tokens in den Quelltext aufnehmen.
2. Vor jedem Push: go build ./... und go vet ./... muessen fehlerfrei durchlaufen.
3. Neue Komponenten unter internal/ anlegen, nicht im Root verstreuen.
4. Aenderungen am Setup-Flow erfordern eine Aktualisierung von web/setup.html und internal/setup/setup.go gemeinsam.

## Agenten

UPHarness ist als Harness fuer autonome Agenten konzipiert. Agenten laufen ueber den Daemon und die Queue. Neue Agententypen werden als Pakete unter internal/agent/ hinzugefuegt und ueber die Konfigurations-API gesteuert.