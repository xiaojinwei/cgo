package service

import (
	"cgo/dao"
	"cgo/entity"
)

type FeedbackService struct {
}

var feedbackDao = new(dao.FeedbackDao)

func (p *FeedbackService) Insert(userId uint,title string,content string) int64 {
	 return feedbackDao.Insert(&entity.Feedback{UserID:userId,Title:title,Content:content})
}
