package main

import (
	"fmt"
	"machine/usb"
	"machine/usb/hid/keyboard"
	"time"
)

func init() {
	usb.Manufacturer = "Two Fat Lads (Heavy Industries)"
	usb.Product = "Bob Hoskins Excuse Boskins"
}

func main() {
	fmt.Println("Booting.....")

	// Trigger some flashes while the USB is recognised
	f := NewFlasher()
	f.FlashyFlash()

	sendString(excuse())

	f.Off()
}

// sendString sends each character in the string as a keyboard input
func sendString(s string) {
	for _, char := range s {
		sendKey(char)
		time.Sleep(10 * time.Millisecond) // Small delay between keystrokes
	}
}

// sendKey maps a character to a keyboard key and sends it over USB HID
func sendKey(char rune) {
	kb := keyboard.New() // Create a new keyboard instance

	var err error

	switch {
	case char >= 97 && char <= 122:
		err = kb.Press(keysToPresses[char]) // Type lowercase letters directly

	case char >= 65 && char <= 90:
		kb.Down(keyboard.KeyModifierLeftShift)
		err = kb.Press(keysToPresses[char+32]) // Type uppercase letters with Shift
		kb.Up(keyboard.KeyModifierLeftShift)

	case char == ' ':
		err = kb.Press(keyboard.KeySpace) // Space character

	case char == ',':
		err = kb.Press(keyboard.KeyComma) // Comma character

	case char == '\'':
		err = kb.Press(keyboard.KeyQuote) // Inverted Comma

	case char == '.':
		err = kb.Press(keyboard.KeyPeriod) // Comma character

	case char == '!':
		kb.Down(keyboard.KeyModifierLeftShift)
		err = kb.Press(keyboard.Key1) // Exclamation mark requires Shift+1
		kb.Up(keyboard.KeyModifierLeftShift)
	}

	if err != nil {
		fmt.Printf("Tried to handle character: %q\n", string(char))

		panic(err)
	}
}
