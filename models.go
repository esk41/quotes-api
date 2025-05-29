package main

import "time"

type Quote struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	Quote     string    `json:"quote"`
	CreatedAT time.Time `json:"created_at"`
}
