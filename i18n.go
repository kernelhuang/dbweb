package main

import (
	"fmt"
	"strings"

	options "github.com/kernelhuang/dbweb/modules/options"
	"github.com/unknwon/i18n"
)

func InitI18n(langs []string) error {
	for _, lang := range langs {
		data, err := options.Locale(fmt.Sprintf("locale_%s.ini", strings.ToLower(lang)))
		if err != nil {
			return err
		}
		i18n.SetMessage(lang, data)
	}
	return i18n.ReloadLangs(langs...)
}
