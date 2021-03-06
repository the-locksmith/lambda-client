package input

import (
	"github.com/galaco/lambda-client/input/keyboard"
	"github.com/galaco/lambda-client/messages"
	"github.com/galaco/lambda-core/event"
	"github.com/galaco/tinygametools"
)

// Keyboard key wrapper
type Keyboard struct {
	keysDown [1024]bool
}

// IsKeyDown Check if a specific key is pressed
func (keyboard *Keyboard) IsKeyDown(key keyboard.Key) bool {
	return keyboard.keysDown[int(key)]
}

// CallbackMouseMove Event manager message receiver.
// Used to catch key events from the window library
func (keyboard *Keyboard) ReceiveMessage(message tinygametools.Event) {
	switch message.Type() {
	case messages.TypeKeyDown:
		keyboard.keysDown[int(message.(*messages.KeyDown).Key)] = true
	case messages.TypeKeyReleased:
		keyboard.keysDown[int(message.(*messages.KeyReleased).Key)] = false
	}
}

func (keyboard *Keyboard) SendMessage() event.IMessage {
	return nil
}

var staticKeyboard Keyboard

// GetKeyboard return static keyboard
func GetKeyboard() *Keyboard {
	return &staticKeyboard
}
