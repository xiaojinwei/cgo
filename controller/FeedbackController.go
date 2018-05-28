package controller

import (
	"cgo/cgo"
	"cgo/service"
	"net/http"
	"cgo/constant"
	"cgo/utils"
)

type FeedbackController struct {
	cgo.ApiController
}

var feedbackService = new(service.FeedbackService)

func (p *FeedbackController)Router(router *cgo.RouterHandler)  {
	router.Router("/feedback",p.feedback)
	router.Router("/getFeedback",p.getFeedback)
}

func (p *FeedbackController)feedback(w http.ResponseWriter,r *http.Request)  {
	userId := p.GetUserId(w,r)
	if userId == 0 {
		cgo.ResultFail(w,"Not logged in")
		return
	}
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	if utils.Empty(title) || utils.Empty(content) {
		cgo.ResultFail(w,"Parameter incomplete")
		return
	}
	to := p.SaveFiles(r,constant.FEEDBACK_IMAGE)
	var images []string
	if to != nil {
		for _,value := range to{
			images = append(images,value.Path)
		}
	}
	id := feedbackService.Insert(userId,title,content,images)
	if id <= 0{
		cgo.ResultFail(w,"feedback fail")
		return
	}
	cgo.ResultOk(w,"feedback success")
}

func (p *FeedbackController)getFeedback(w http.ResponseWriter,r *http.Request)  {
	userId := p.GetUserId(w,r)
	if userId == 0 {
		cgo.ResultFail(w,"Not logged in")
		return
	}
	fd := feedbackService.SelectFeedbackByUserId(userId)
	cgo.ResultJsonOk(w,fd)
}
