package models

import "time"

type User struct {
	ID        int64     `db:"id"`
	Phone     string    `db:"phone"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type OTPData struct {
	Code     int32  `json:"code"`
	Phone    string `json:"phone"`
	Tries    int    `json:"tries"`
}
