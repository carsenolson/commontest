package controllers

import (
	"net/http"
)

func Index(rw http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(rw, "index.html", nil)
}
