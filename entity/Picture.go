package entity

import "time"

type Picture struct {
	ID uint `json:"id"`
	FeedbackID uint `json:"feedback_id"`
	Address string `json:"address"`
	CreateTime time.Time `json:"create_time"`
} 

