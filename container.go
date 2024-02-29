package state_manager

// The Container operates with values and provides an API for reading, deletion, etc
type Container interface {
	// Get returns the value by key and additionally a boolean flag
	// indicating whether the container contains a value for that key
	Get(key ValueKey) (Value, bool)
	// Read returns the value by key
	Read(key ValueKey) Value
	// Exists returns a boolean indicating whether the key exists or not
	Exists(key ValueKey) bool
	// Add adds the value to the container
	// Be careful: it doesn't check if the key already exists, so it will override the existing value
	Add(key ValueKey) error
	// Remove removes the value from the storage.
	// If the key doesn't exist, it won't return an error regarding that
	Remove(key ValueKey) error
}
