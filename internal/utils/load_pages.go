package utils

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageTamplate map[string]*template.Template

// templates is where the page html templates stored
var templates PageTamplate

// LoadTemplates loads all the html from web/pages
func LoadTemplates() {
	templates = make(PageTamplate)

	//tPages list all the pages that used layout.html
	tPages := []string{"dashboard", "transaction"}
	//uPages list all the pages that has their own template
	uPages := []string{"login", "404"}

	for _, page := range tPages {
		t := template.Must(template.ParseFiles(
			"web/pages/layout.html",
			"web/pages/"+page+"/"+page+".html",
		))
		templates[page] = t
	}

	for _, page := range uPages {
		t := template.Must(template.ParseFiles(
			"web/pages/" + page + "/" + page + ".html",
		))
		templates[page] = t
	}
}

// Render render the web page from templates
func Render(c *gin.Context, page string, data gin.H) {
	t, ok := templates[page]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "page not found"})
		return
	}

	c.Status(http.StatusOK)
	c.Header("Content-Type", "text/html; charset=utf-8")
	t.ExecuteTemplate(c.Writer, "layout", data)
}
