package actions

import (
	"fmt"

	"github.com/kernelhuang/binding"
	"github.com/kernelhuang/dbweb/models"
	"github.com/kernelhuang/flash"
	"github.com/kernelhuang/renders"
	"github.com/kernelhuang/xsrf"
	"github.com/unknwon/i18n"
)

type Addb struct {
	AuthRenderBase

	binding.Binder
	xsrf.Checker
	flash.Flash
}

func (c *Addb) Get() error {
	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("add.html", renders.T{
		"dbs":          SupportDBs,
		"flash":        c.Flash.Data(),
		"engines":      engines,
		"XsrfFormHtml": c.XsrfFormHtml(),
		"IsLogin":      c.IsLogin(),
		"isAdd":        true,
	})
}

func (c *Addb) Post() {
	var engine models.Engine
	engine.Name = c.Form("name")
	engine.Driver = c.Form("driver")
	host := c.Form("host")
	port := c.Form("port")
	dbname := c.Form("dbname")
	username := c.Form("username")
	passwd := c.Form("passwd")

	if engine.Driver == "sqlite3" {
		engine.DataSource = host
	} else {
		engine.DataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			username, passwd, host, port, dbname)
	}

	/*if err := c.MapForm(&engine); err != nil {
		c.Flash.Set("ErrAdd", i18n.Tr(c.CurLang(), "err_param"))
		c.Redirect("/addb")
		return
	}*/

	if err := models.AddEngine(&engine); err != nil {
		c.Flash.Set("ErrAdd", i18n.Tr(c.CurLang(), "err_add_failed"))
		c.Redirect("/addb")
		return
	}

	c.Redirect("/")
}
