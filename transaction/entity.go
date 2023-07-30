package transaction

import (
	"startup-funding/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserId     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
