package entities

import (
	"time"
)

const (
	DefaultTeg TegName = "default"
	AnimalTeg  TegName = "animals"
	SpaceTeg   TegName = "space"
)

type TegName string

type FormData struct {
	Id        string        `json:"id" db:"id"`
	Title     string        `json:"title" db:"title"`
	Text      string        `json:"text" db:"text"`
	Tag       TegName       `json:"tag" db:"tag"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
