package values

import (
	"sync"
	"time"
)

type UsersByIdsValueItemsRequest struct {
	ID uint
}

const StateValuesUsersByIdsKey = "users_by_ids"

// UsersByIdsValue
// This contains a map of active users with their IDs and email addresses,
// where the key is the user ID and the value is the complete information about the user
type UsersByIdsValue struct {
	// Key - id
	// Value - [User]
	users sync.Map
}

func (u *UsersByIdsValue) Update() error {
	// Any updates, whether from storage (such as a database), an API, random generation, etc.

	return nil
}

func (u *UsersByIdsValue) UpdateSchedule() time.Duration {
	return 0
}

func (u *UsersByIdsValue) UpdateTTL() time.Duration {
	return 24 * time.Hour
}

func (u *UsersByIdsValue) Items(requests ...any) (any, error) {
	if len(requests) == 0 {
		return u.users, nil
	}

	req, ok := requests[0].(*UsersByIdsValueItemsRequest)
	if ok {
		user, _ := u.users.Load(req.ID)

		return user, nil
	}

	return nil, nil
}
