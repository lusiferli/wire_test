package domain

import (
	"fmt"
	"git.code.oa.com/tc-framework-go/errs"
	"github.com/looplab/fsm"
)

type Package struct {
	Id        int
	statusFms *fsm.FSM
}

var (
	UnSubmit  = "UnSubmit"
	Submitted = "Submitted"
	Reviewed  = "Reviewed"
)

var (
	EventSubmit = "submit"
	EventReview = "review"
)

var (
	Unknown         = 0
	StatusUnSubmit  = 1
	StatusSubmitted = 2
	StatusReviewed  = 3
)

func ToStatusCode(status string) int {
	statusCode := 0
	switch status {
	case UnSubmit:
		statusCode = StatusUnSubmit
	case Submitted:
		statusCode = StatusSubmitted
	case Reviewed:
		statusCode = StatusReviewed
	default:
		statusCode = Unknown
	}
	return statusCode
}

func atterEvent(event string) string {
	return "after_" + event
}

type SubmitCommand struct {
}

var CodePackageStatusChange = 500001

func (p *Package) Submit(cmd SubmitCommand) error {
	if !p.statusFms.Can(EventSubmit) {
		currentStatus := p.statusFms.Current()
		return errs.NewIllegalStateError("当前状态是 : "+currentStatus+", 不允许提交", nil)
	}
	err := p.statusFms.Event(EventSubmit)
	if err != nil {
		return errs.NewDomainError(CodePackageStatusChange, "当前包状态未能发生变化", nil)
	}
	return nil
}

func (p *Package) Status() int {
	return ToStatusCode(p.statusFms.Current())
}

func NewPackage() *Package {
	p := Package{}

	p.statusFms = fsm.NewFSM(
		// 初始状态未提交
		UnSubmit,
		fsm.Events{
			// 提交事件
			{Name: EventSubmit, Src: []string{UnSubmit}, Dst: Submitted},
			// review 事件
			{Name: EventReview, Src: []string{Submitted}, Dst: Reviewed},
		},
		fsm.Callbacks{
			atterEvent(EventSubmit): func(e *fsm.Event) {
				fmt.Println("after event submit")
			},
			atterEvent(EventReview): func(e *fsm.Event) {
				fmt.Println("after event review")
			},
		},
	)
	return &p
}
