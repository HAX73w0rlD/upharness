package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Bridge ist eine Telegram-Adapter mit Long-Polling.
// Token kommt NUR aus der Umgebung, niemals hardcodiert.
type Bridge struct {
	Token  string
	Active bool
	offset int64
	client *http.Client
}

// Update ist eine Telegram-Update-Nachricht.
type Update struct {
	UpdateID int64  `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

// Message ist eine Telegram-Nachricht.
type Message struct {
	MessageID int64  `json:"message_id"`
	From      *User  `json:"from,omitempty"`
	Chat      *Chat  `json:"chat,omitempty"`
	Text      string `json:"text,omitempty"`
	Date      int64  `json:"date"`
}

// User ist ein Telegram-Benutzer.
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
}

// Chat ist ein Telegram-Chat.
type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type,omitempty"`
	Title     string `json:"title,omitempty"`
}

// NewFromEnv erstellt eine Bridge aus Umgebungsvariablen.
func NewFromEnv() *Bridge {
	return &Bridge{Token: os.Getenv("TELEGRAM_BOT_TOKEN")}
}

// Start beginnt das Long-Polling.
func (b *Bridge) Start() error {
	if b.Token == "" && os.Getenv("TELEGRAM_DRYRUN") != "1" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN nicht gesetzt")
	}
	b.Active = true
	b.client = &http.Client{Timeout: 30 * time.Second}
	go b.pollLoop()
	return nil
}

// Stop beendet die Bridge.
func (b *Bridge) Stop() { b.Active = false }

// pollLoop lauf das Long-Polling gegen die Telegram API.
func (b *Bridge) pollLoop() {
	for b.Active {
		if err := b.pollOnce(); err != nil {
			fmt.Println("Telegram Poll-Fehler:", err)
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

// pollOnce fuhrt einen einzelnen Poll-Zyklus aus.
func (b *Bridge) pollOnce() error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d&timeout=25", b.Token, b.offset+1)
	resp, err := b.client.Get(url)
	if err != nil { return err }
	defer resp.Body.Close();
	body, err := io.ReadAll(resp.Body)
	if err != nil { return err }
	var result struct {
		OK     bool    `json:"ok"`
		Result []Update `json:"result"`
	}
	if err := json.Unmarshal(body, &result); err != nil { return err }
	if !result.OK { return fmt.Errorf("Telegram API nicht ok") }
	for _, upd := range result.Result {
		b.offset = upd.UpdateID
		if upd.Message != nil {
			b.handleMessage(upd.Message)
		}
	}
	return nil
}

// handleMessage verarbeitet eine eingehende Nachricht.
func (b *Bridge) handleMessage(m *Message) {
	if m.Text == "" { return }
	text := strings.TrimSpace(m.Text)
	if strings.HasPrefix(text, "/start") {
		b.sendText(m.Chat.ID, "Willkommen bei UPHarness. Sende mir eine Nachricht.")
		return
	}
	if strings.HasPrefix(text, "/help") {
		b.sendText(m.Chat.ID, "Verfugbare Befehle: /start, /help, /status")
		return
	}
	if strings.HasPrefix(text, "/status") {
		b.sendText(m.Chat.ID, "UPHarness lauft. Aktive Sessions: laufend.")
		return
	}
	// Echo als Platzhalter für echte KI-Integration
	b.sendText(m.Chat.ID, "Empfangen: "+text)
}

// sendText sendet eine Textnachricht an einen Chat.
func (b *Bridge) sendText(chatID int64, text string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", b.Token, chatID, urlEncode(text))
	resp, err := b.client.Get(url)
	if err != nil { fmt.Println("Telegram Send-Fehler:", err); return }
	defer resp.Body.Close();
}

// urlEncode kodiert einen String fuer URLs.
func urlEncode(s string) string {
	return strings.NewReplacer(" ", "%20", "\n", "%0A", "&", "%26", "=", "%3D").Replace(s)
}
