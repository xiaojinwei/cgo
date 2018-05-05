package dao

import (
	"cgo/entity"
	"cgo/cgo"
	"log"
)

type FeedbackDao struct {
}

func (p *FeedbackDao) Insert(feedback *entity.Feedback) int64 {
	result,err := cgo.DB.Exec("INSERT INTO feedback(`user_id`,`title`,`content`,`create_time`) VALUE(?,?,?,?)",feedback.UserID,feedback.Title,feedback.Content,feedback.CreateTime)
	if err != nil {
		log.Println(err)
		return 0
	}
	id,err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}
