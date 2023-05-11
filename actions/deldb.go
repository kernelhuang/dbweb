package actions

import (
	"strconv"

	"github.com/kernelhuang/dbweb/models"
	"github.com/kernelhuang/renders"
)

type Del struct {
	AuthRenderBase
}

func (c *Del) Get() error {
	id, err := strconv.ParseInt(c.Req().FormValue("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := models.DelEngineById(id); err != nil {
		return err
	}

	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("delsuccess.html", renders.T{
		"engines": engines,
		"IsLogin": c.IsLogin(),
	})
}
