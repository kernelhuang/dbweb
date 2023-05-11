package actions

import (
	"github.com/kernelhuang/flash"
	"github.com/kernelhuang/renders"
	"github.com/kernelhuang/xsrf"
	"github.com/unknwon/i18n"

	"github.com/kernelhuang/dbweb/models"
)

type ChgPass struct {
	AuthRenderBase

	xsrf.Checker
	flash.Flash
}

func (c *ChgPass) Get() error {
	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("chgpass.html", renders.T{
		"XsrfFormHtml": c.XsrfFormHtml(),
		"engines":      engines,
		"flash":        c.Flash.Data(),
		"IsLogin":      c.IsLogin(),
	})
}

func (c *ChgPass) Post() {
	oldPass := c.Req().FormValue("old_pass")
	newPass := c.Req().FormValue("new_pass")
	cfmPass := c.Req().FormValue("cfm_pass")

	defer c.Redirect("/chgpass")

	if newPass != cfmPass {
		c.Flash.Set("cfmError", i18n.Tr(c.CurLang(), "password_not_eq"))
		return
	}

	user := c.LoginUser()
	if user != nil {
		if models.EncodePassword(oldPass) != user.Password {
			c.Flash.Set("oldError", i18n.Tr(c.CurLang(), "ori_password_not_correct"))
			return
		}
	} else {
		c.Flash.Set("otherError", i18n.Tr(c.CurLang(), "unknown_error"))
		return
	}

	user.Password = newPass
	err := models.UpdateUser(user)
	if err != nil {
		c.Flash.Set("otherError", err.Error())
		return
	}

	c.Flash.Set("changeSuccess", i18n.Tr(c.CurLang(), "password_change_success"))
}
