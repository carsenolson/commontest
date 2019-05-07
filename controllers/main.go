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
		// Don't mind lol
		"answerType": func(index int, right, picked []int) template.HTMLAttr {
			required := compare(right, picked)
			additional := compare(picked, right)
			//rgb(248, 249, 250) nothing
			//rgb(141, 212, 157) possible right answer	
			//rgb(40, 167, 69) picked answer
			//rgb(220, 53, 69) mistake
			for _, req := range required {
				if index == req {
					return template.HTMLAttr(`style="background-color: rgb(141, 212, 157);"`)
				}
			}
			for _, add := range additional {
				if index == add {
					return template.HTMLAttr(`style="background-color: rgb(220, 53, 69);"`)
				}
			}
			right_picked_answers := compare(right, required)
			for _, rpa := range right_picked_answers {
				if index == rpa {
					return template.HTMLAttr(`style="background-color: rgb(40, 167, 69);"`)
				}
			}
			return template.HTMLAttr(`style="background-color: rgb(248, 249, 250);"`)
		},
	}
	tpl = template.Must(template.New("Template").Funcs(funcMap).ParseGlob("templates/*.html"))
}
