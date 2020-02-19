package pkg

// GuiMessage represents a message to be sent or receibed from the UI
type GuiMessage struct {
	EventType byte
}

// MessageTypes Represents all the event types that GBGO can send to the UI
var MessageTypes = struct {
	MemoryChanged byte
	NewFrame      byte
}{
	MemoryChanged: 0,
	NewFrame:      1,
}
