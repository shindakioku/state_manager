package state_manager

import "time"

// ValueKey is the unique identifier of a value
type ValueKey string

// The Value contract represents each value for the state
// It provides configuration for the state manager and persists the user's data
type Value interface {
	// Update is the function called periodically or manually to perform updates on the internal state
	// For instance, in some implementations, a value may contain a collection,
	// each update call should update that state, if possible
	Update() error
	// UpdateSchedule specifies a time period for the periodic updates
	UpdateSchedule() time.Duration
	// Items provides an API for reading the internal state
	// Additionally, it can receive any arguments that are be able to used as filters
	// Please take a look at [examples/values/users_by_ids.go]
	Items(requests ...any) (any, error)
	// Stop the function calls when [State] is triggered by a shutdown signal
	// You ought to close any connections, destroy objects, etc
	Stop() error
}
