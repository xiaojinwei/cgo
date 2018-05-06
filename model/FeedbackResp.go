package model

import "cgo/entity"

type FeedbackResp struct {
	entity.Feedback
	Pictures []entity.Picture
}

