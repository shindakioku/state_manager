package values

import (
	"state_manager"
	"time"
)

const StateValuesUsersKey = "users"

// UsersValue
// This contains a list of active users with their IDs and email addresses
type UsersValue struct {
	// @PLEASE DO NOT USE THIS EXAMPLE IN PRODUCTION!
	// SLICES ARE NOT CONCURRENTLY SAFE
	users []User
}

func (u *UsersValue) Update() error {
	// Any updates, whether from storage (such as a database), an API, random generation, etc.

	return nil
}

func (u *UsersValue) Items(requests ...any) (any, error) {
	return u.users, nil
}

func (u *UsersValue) UpdateSchedule() time.Duration {
	return 1 * time.Hour
}

// NewUsersValue
// You can provide any data you want, but please don't initialize the state.
// It's not useful because the state manager will immediately call [Update] on the first run.
func NewUsersValue() state_manager.Value {
	return &UsersValue{}
}
