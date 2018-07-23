package dao

import (
	"cgo/entity"
	"cgo/cgo"
	"log"
	"cgo/model"
	"cgo/dao/bean"
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

func (p *FeedbackDao) SelectFeedbackByUserId(id uint) []*model.FeedbackResp {
	rows,err := cgo.DB.Query("SELECT * FROM feedback f,picture p WHERE f.user_id = ? AND f.id = p.feedback_id",id)
	if err != nil {
		log.Println(err)
		return nil
	}
	var resps []*model.FeedbackResp
	for rows.Next() {
		var temp bean.TempFeedback
		err := rows.Scan(&temp.Feedback.ID,&temp.Feedback.UserID,&temp.Feedback.Title,&temp.Feedback.Content,&temp.Feedback.CreateTime,
			&temp.Picture.ID,&temp.Picture.FeedbackID,&temp.Picture.Address,&temp.Picture.CreateTime)
		if err != nil{
			log.Println(err)
			continue
		}
		resps = p.appendFeedback(resps,&temp)
	}
	rows.Close()
	return resps
}

func (p *FeedbackDao) appendFeedback(resps []*model.FeedbackResp,temp *bean.TempFeedback) []*model.FeedbackResp {
	var id uint
	if len(resps) > 0 {
		id = resps[len(resps) - 1].Feedback.ID
	}
	if id == temp.Feedback.ID {
		picture := entity.Picture{temp.Picture.ID,temp.Picture.FeedbackID,temp.Picture.Address,temp.Picture.CreateTime}
		//取切片
		pictures := resps[len(resps) - 1].Pictures
		pictures = append(pictures,picture)
		//重新赋值切片
		resps[len(resps) - 1].Pictures = pictures
	} else {
		feedback := entity.Feedback{temp.Feedback.ID,temp.Feedback.UserID,temp.Feedback.Title,temp.Feedback.Content,temp.Feedback.CreateTime}
		picture := entity.Picture{temp.Picture.ID,temp.Picture.FeedbackID,temp.Picture.Address,temp.Picture.CreateTime}
		resp := model.FeedbackResp{Feedback:feedback}
		resp.Pictures = append(resp.Pictures,picture)
		resps = append(resps,&resp)
	}
	return resps
}

































