package actions

import (
	"github.com/kernelhuang/dbweb/middlewares"
	"github.com/lunny/tango"
)

type Logout struct {
	RenderBase
	middlewares.AuthUser
	tango.Ctx
}

func (l *Logout) Get() {
	if l.IsLogin() {
		l.Logout()
	}
	l.Redirect("/")
}
