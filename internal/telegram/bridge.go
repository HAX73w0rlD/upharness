package telegram

import "os"

// Bridge is a slim Telegram adapter. Token comes ONLY from env, never hardcoded.
// Env: TELEGRAM_BOT_TOKEN (secret, local only), TELEGRAM_DRYRUN=1 for tests.
type Bridge struct {
	Token string
	Active bool
	queue []string
}

func NewFromEnv() *Bridge { return &Bridge{Token: os.Getenv("TELEGRAM_BOT_TOKEN")} }

// Enqueue holds a message while a job runs. Dequeue must run on every exit path.
func (b *Bridge) Enqueue(text string) { b.queue = append(b.queue, text) }

func (b *Bridge) Dequeue() (string, bool) {
	if len(b.queue)==0 { return "", false }
	t := b.queue[0]; b.queue = b.queue[1:]; return t, true
}

func (b *Bridge) Start() error { b.Active = b.Token != "" || os.Getenv("TELEGRAM_DRYRUN") == "1"; return nil }
