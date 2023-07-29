package transaction

import "time"

type Transaction struct {
	ID         int
	CampaignID int
	UserId     int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
