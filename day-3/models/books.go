package models

import (
	"time"
)

type Book struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title" validate:"required"`
	Author      string    `json:"author" validate:"required"`
	PublishDate time.Time `json:"publish_date"`
}
