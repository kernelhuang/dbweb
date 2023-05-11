package main

import (
	"html/template"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/kernelhuang/dbweb/modules/public"
	"github.com/kernelhuang/dbweb/modules/templates"

	"github.com/kernelhuang/binding"
	"github.com/kernelhuang/captcha"
	"github.com/kernelhuang/dbweb/actions"
	"github.com/kernelhuang/dbweb/middlewares"
	"github.com/kernelhuang/debug"
	"github.com/kernelhuang/flash"
	"github.com/kernelhuang/renders"
	"github.com/kernelhuang/session"
	"github.com/lunny/nodb"
	"github.com/lunny/tango"
	"github.com/unknwon/i18n"
)

var (
	sessionTimeout = time.Minute * 20
)

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	aa := reflect.ValueOf(a)
	return !aa.IsValid() || (aa.Type().Kind() == reflect.Ptr && aa.IsNil())
}

func InitTango(isDebug bool) *tango.Tango {
	xormVersion := "v1.3.2"
	t := tango.New()
	if isDebug {
		t.Use(debug.Debug(debug.Options{
			HideResponseBody: true,
			IgnorePrefix:     "/public",
		}))
	}
	sess := session.New(session.Options{
		MaxAge: sessionTimeout,
	})
	t.Use(
		tango.Logging(),
		tango.Recovery(false),
		tango.Compresses([]string{}),
		public.Static(),
		tango.Return(),
		tango.Param(),
		tango.Contexts(),
		binding.Bind(),
		renders.New(renders.Options{
			Reload:    true,
			Directory: "templates",
			Funcs: template.FuncMap{
				"isempty": func(s string) bool {
					return len(s) == 0
				},
				"add": func(a, b int) int {
					return a + b
				},
				"isNil": isNil,
				"i18n":  i18n.Tr,
				"Range": func(size int) []struct{} {
					return make([]struct{}, size)
				},
				"multi": func(a, b int) int {
					return a * b
				},
			},
			Vars: renders.T{
				"GoVer":    strings.Trim(runtime.Version(), "go"),
				"TangoVer": tango.Version(),
				"XormVer":  xormVersion,
				"NodbVer":  nodb.Version,
			},
			FileSystem: templates.FileSystem("templates"),
		}),
		captcha.New(),
		sess,
		middlewares.Auth("/login"),
		flash.Flashes(sess),
	)

	t.Any("/", new(actions.Home))
	t.Any("/login", new(actions.Login))
	t.Any("/logout", new(actions.Logout))
	t.Any("/addb", new(actions.Addb))
	t.Any("/view", new(actions.View))
	t.Any("/del", new(actions.Del))
	t.Any("/delRecord", new(actions.DelRecord))
	t.Any("/chgpass", new(actions.ChgPass))
	t.Get("/test", new(actions.Test))
	return t
}
