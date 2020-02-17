package main

// GuiMessage represents a message to be sent or receibed from the UI
type GuiMessage struct {
	EventType byte
}

// MessageTypes Represents all the event types that GBGO can send to the UI
var MessageTypes = struct {
	memoryChanged byte
	newFrame      byte
}{
	memoryChanged: 0,
	newFrame:      1,
}
