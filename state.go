package state_manager

// The State contract is the primary interface that provides you with an API
// to access persisted values (via the [Container] contract) and allows for updates (manually/periodically).
type State interface {
	// Values returns an instance of the container that persists values
	Values() Container
	// @TODO
	Updates() error
	// @TODO
	ManuallyUpdates(keys ...ValueKey) error
	// @TODO
	Stop() error
}
