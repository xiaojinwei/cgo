package controller

import (
	"cgo/cgo"
	"cgo/service"
	"net/http"
)

type FeedbackController struct {
	cgo.Controller
}

var feedbackService = new(service.FeedbackService)

func (p *FeedbackController)Router(router *cgo.RouterHandler)  {
	router.Router("/feedback",p.feedback)
}

func (p *FeedbackController)feedback(w http.ResponseWriter,r *http.Request)  {

}
