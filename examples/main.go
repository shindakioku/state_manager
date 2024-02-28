package main

import (
	"log"
	"state_manager/examples/values"
)

func main() {
	// Create instances of your values
	usersValue := values.NewUsersValue()
	usersByIdsValue := state_manager.Value(&values.UsersByIdsValue{})

	// Create an instance of the state manager and pass your values
	stateManager := state_manager.New(usersValue, usersByIdsValue)
	// Don't forget to call the updates handler, which will trigger periodic updates of your values.
	go stateManager.Updates()

	// Utilize your values
	// Generic version
	users := state_manager.Get[[]values.User](stateManager, values.StateValuesUsersKey)
	// Type assertion version
	users := stateManager.Get(values.StateValuesUsersKey).([]values.User)
	// User by id
	user := stateManager.Get(values.StateValuesUsersByIdsKey, values.UsersByIdsValueItemsRequest{ID: 1}).(values.User)

	// Opportunities
	// Manually trigger the updater for each value in the storage
	if err := stateManager.ManuallyUpdates(); err != nil {
		// The first error occurred
	}

	// Manually trigger the updater for only certain values
	if err := stateManager.ManuallyUpdates(values.StateValuesUsersKey, ...); err != nil {
		// The first error occurred
	}

	// Retrieve a list of all stored values.
	values := stateManager.Values()
}
