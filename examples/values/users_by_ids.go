package values

import "sync"

const StateValuesUsersByIdsKey = "users_by_ids"

type UsersByIdsValue struct {
	// Key - id
	// Value - [User]
	users sync.Map
}

// Key Unique key
func (u *UsersByIdsValue) Key() state_manager.ValueKey {
	return StateValuesUsersByIdsKey
}

func (u *UsersByIdsValue) Description() string {
	return `
		This contains a map of active users with their IDs and email addresses, 
		where the key is the user ID and the value is the complete information about the user
	`
}

func (u *UsersByIdsValue) Update() error {
	// Any updates, whether from storage (such as a database), an API, random generation, etc.

	return nil
}

func (u *UsersByIdsValue) Items() any {
	return u.users
}
