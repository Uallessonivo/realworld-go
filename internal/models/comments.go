package models

import "time"

type Comment struct {
	id        int
	createdAt time.Time
	updatedAt time.Time
	body      string
	author    Profile
}
