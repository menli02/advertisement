package models

import "time"

type Advertisement struct {
	ID          int64      `db:"id"`
	UserID      int64      `db:"user_id"`
	CategoryID  int64      `db:"category_id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Slug        string     `db:"slug"`
	Price       float64    `db:"price"`
	Currency    string     `db:"currency"`
	Status      string     `db:"status"`
	ViewCount   int64      `db:"view_count"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

type Category struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Slug      string    `db:"slug"`
	CreatedAt time.Time `db:"created_at"`
}

type AdImage struct {
	ID        int64     `db:"id"`
	AdID      int64     `db:"ad_id"`
	URL       string    `db:"url"`
	Position  int       `db:"position"`
	CreatedAt time.Time `db:"created_at"`
}
