package controllers

import (
	"html/template"
)

var tpl *template.Template

func init() {
	funcMap := template.FuncMap{
		"safeUrl": func(s string) template.URL {
			return template.URL(s)
		},
		"isInRightAnswers": func(num int, right []int) bool {
			for _ ,ra := range right {
				if num == ra {
					return true
				}
				continue
			}
			return false
		},
	}
	tpl = template.Must(template.New("Template").Funcs(funcMap).ParseGlob("templates/*.html"))
}
