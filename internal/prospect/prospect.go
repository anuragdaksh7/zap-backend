package prospect

import "time"

type NewProspect struct {
	FullName    string     `json:"full_name"`
	Email       string     `json:"email"`
	Phone       *string    `json:"phone"`
	Company     *string    `json:"company"`
	Position    *string    `json:"position"`
	LinkedInURL *string    `json:"linked_in_url"`
	Status      *string    `json:"status"`
	Location    *string    `json:"location"`
	LastContact *time.Time `json:"last_contact"`
	ResponseAt  *time.Time `json:"response_at"`
	Notes       *string    `json:"notes"`
}
