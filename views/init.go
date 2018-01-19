package views

import (
	"html/template"
	"utils"
)

var (
	//TplFuncs contain the functions will be used in tpl file
	TplFuncs = make(template.FuncMap)
)

func init() {
	TplFuncs["TimeNow"] = utils.TimeNow
	TplFuncs["FormatTime"] = utils.FormatTime
}
