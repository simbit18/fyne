package fyne

// ScanCode represents a hardware ID for (normally desktop) keyboard events.
// Most applications should use KeyName for cross-platform compatibility.
type ScanCode int

// KeyEvent describes a keyboard input event.
type KeyEvent struct {
	// Name describes the keyboard event that is consistent across platforms.
	Name KeyName
	// HardwareCode is a platform specific field that reports the hardware ID of the keyboard events.
	HardwareCode ScanCode
}

// PointEvent describes a pointer input event. The position is relative to the
// top-left of the CanvasObject this is triggered on.
type PointEvent struct {
	AbsolutePosition Position // The absolute position of the event
	Position         Position // The relative position of the event
}

// ScrollEvent defines the parameters of a pointer or other scroll event.
// The DeltaX and DeltaY represent how large the scroll was in two dimensions.
type ScrollEvent struct {
	PointEvent
	Scrolled Delta
}

// DragEvent defines the parameters of a pointer or other drag event.
// The DraggedX and DraggedY fields show how far the item was dragged since the last event.
type DragEvent struct {
	PointEvent
	Dragged Delta
}
