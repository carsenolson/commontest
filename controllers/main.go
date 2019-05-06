package controllers

import (
	"fmt"
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
			fmt.Println("right, picked:", right, picked)
			required := compare(right, picked)
			additional := compare(picked, right)
			fmt.Println("required, additional:", required, additional)
			if len(right) == 0 {
				return template.HTMLAttr(`style="background-color: rgb(248, 249, 250);"`)
			}
			if len(picked) == 0 {
				for _, rt := range right {
					if index == rt {
						return template.HTMLAttr(`style="background-color: rgb(40, 167, 69);"`)
					}
				}
			}
			if len(required) != 0 {
				for _, rt := range right {
					if index == rt {
						for pt := range picked {
							if index == pt {
								return template.HTMLAttr(`style="background-color: rgb(40, 167, 69);"`)
							}
						}
                    }
				}
				for _, req := range required {
					if index == req {
						return template.HTMLAttr(`style="background-color: rgb(130, 234, 153);"`)
					}
				}
			}
			if len(additional) != 0 {
					for _, add := range additional {
						if index == add {
							return template.HTMLAttr(`style="background-color: rgb(220, 53, 69);"`)
						}
					}
			}
			if len(required) == 0 && len(additional) == 0 {
				for _, rt := range right {
					if index == rt {
						return template.HTMLAttr(`style="background-color: rgb(40, 167, 69);"`)
					}
				}
			}
			return template.HTMLAttr(`style="background-color: rgb(248, 249, 250);"`)
		},
	}
	tpl = template.Must(template.New("Template").Funcs(funcMap).ParseGlob("templates/*.html"))
}
