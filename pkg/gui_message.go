package pkg

// GuiMessage represents a message to be sent or receibed from the UI
type GuiMessage struct {
	EventType byte
}

const (
	MemoryUpdated = iota
	NewFrame      = iota
	Log           = iota
)
