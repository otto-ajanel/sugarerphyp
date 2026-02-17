package core

import "time"

type User struct {
	ID        int       `json:"id_user"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UserType  int       `json:"id_usertype"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
	TenantID  int       `json:"id_tenant"`
}

// Domain event for user retrieved
type UserRetrievedEvent struct {
	User User
}
