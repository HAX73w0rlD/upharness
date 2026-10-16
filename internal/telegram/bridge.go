package telegram

// Bridge is a lightweight Telegram bridge stub.
type Bridge struct {
	Token  string
	Active bool
}

// NewBridge creates a new Telegram bridge.
func NewBridge(token string) *Bridge {
	return &Bridge{Token: token, Active: token != ""}
}

// Start begins polling (stub).
func (b *Bridge) Start() error {
	b.Active = true
	return nil
}
§uçâçn×ŸŠw