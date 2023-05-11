package actions

import (
	"github.com/kernelhuang/dbweb/models"
	"github.com/kernelhuang/renders"
	"xorm.io/core"
)

type Home struct {
	AuthRenderBase
}

func (c *Home) Get() error {
	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("root.html", renders.T{
		"engines": engines,
		"tables":  []core.Table{},
		"records": [][]string{},
		"columns": []string{},
		"id":      0,
		"ishome":  true,
		"IsLogin": c.IsLogin(),
	})
}
