package store

import (
	"fmt"
	"time"
)

type User struct {
	ID             int64      `json:"id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone"`
	CountryCode    string     `json:"country_code"`
	PasswordHash   string     `json:"-"`
	CreatedAt      time.Time  `json:"created_at"`
	LastSignedIn   *time.Time `json:"last_signed_in"`
	LastSignedInIP *string    `json:"last_signed_in_ip"`
	ArchivedAt     *time.Time `jsob:"archived_at"`
}

type UserStore interface {
	Create(user User) (int64, error)
	Update(user User) error
	Deactivate(id int64) error
	Get(id int64) (*User, error)
	Count() (int64, error)
}

func (u *User) getDisplayName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
