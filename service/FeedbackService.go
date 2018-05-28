package service

import (
	"cgo/dao"
	"cgo/entity"
	"time"
	"cgo/model"
)

type FeedbackService struct {
}

var feedbackDao = new(dao.FeedbackDao)
var pictureDao = new(dao.PictureDao)

func (p *FeedbackService) Insert(userId uint,title string,content string,pictures []string) int64 {
	id := feedbackDao.Insert(&entity.Feedback{UserID:userId,Title:title,Content:content,CreateTime:time.Now()})
	if id <= 0 {
		return 0
	}
	for _,value := range pictures{
		pictureDao.Insert(&entity.Picture{FeedbackID:uint(id),Address:value,CreateTime:time.Now()})
	}
	return id
}

func (p *FeedbackService)SelectFeedbackByUserId(id uint) []*model.FeedbackResp {
	return feedbackDao.SelectFeedbackByUserId(id)
}