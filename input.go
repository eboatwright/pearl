package pearl


import (
	"github.com/hajimehoshi/ebiten/v2"
)


var (
	keys             [110]bool
	keysLast         [110]bool
	mouseButtons     [3]bool
	mouseButtonsLast [3]bool
)


func updateInput() {
	keysLast = keys
	for i := 0; i < len(keys); i++ {
		keys[i] = ebiten.IsKeyPressed(ebiten.Key(i))
	}

	mouseButtonsLast = mouseButtons
	for i := 0; i < len(mouseButtons); i++ {
		mouseButtons[i] = ebiten.IsMouseButtonPressed(ebiten.MouseButton(i))
	}
}

// Returns true or false when key is being held
func KeyPressed(key ebiten.Key) bool {
	return keys[int(key)]
}

// Returns true or false if key was just pressed this frame
func KeyJustPressed(key ebiten.Key) bool {
	return keys[int(key)] && !keysLast[int(key)]
}

// Returns true or false if key was just released this frame
func KeyJustReleased(key ebiten.Key) bool {
	return !keys[int(key)] && keysLast[int(key)]
}

// Returns true or false when mouse button is being held
func MouseButtonPressed(button ebiten.MouseButton) bool {
	return mouseButtons[int(button)]
}

// Returns true or false if mouse button was just pressed this frame
func MouseButtonJustPressed(button ebiten.MouseButton) bool {
	return mouseButtons[int(button)] && !mouseButtonsLast[int(button)]
}

// Returns true or false if mouse button was just released this frame
func MouseButtonJustReleased(button ebiten.MouseButton) bool {
	return !mouseButtons[int(button)] && mouseButtonsLast[int(button)]
}

// Get mouse position as Vector2
func GetMousePosition() Vector2 {
	x, y := ebiten.CursorPosition()
	return Vector2 { float64(x), float64(y) }
}


// Returns an integer and is -1 if one of the negative keys specified is pressed, and returns 1 if any of the positive keys specified is pressed. Returns 0 if both or none are pressed.
func GetInputAxis(negative, positive []ebiten.Key) int {
	result := 0

	for _, key := range negative {
		if KeyPressed(key) {
			result -= 1
			break
		}
	}

	for _, key := range positive {
		if KeyPressed(key) {
			result += 1
			break
		}
	}

	return result
}