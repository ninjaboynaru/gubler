package server

import (
	"bytes"
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

func createTemplate(templateName string) *template.Template {
	return template.Must(template.ParseFiles("static/html/layout.tmpl", "static/html/"+templateName+".tmpl", "static/html/header.tmpl", "static/html/footer.tmpl"))
}

var staticTemplateMap = map[string]*template.Template{
	"posts":      createTemplate("posts"),
	"createpost": createTemplate("createpost"),
	"about":      createTemplate("about"),
}

func executeTemplateStatic(templateName string, buffer *bytes.Buffer) {
	staticTemplateMap[templateName].ExecuteTemplate(buffer, "layout", nil)
}

func executeTemplateDynamic(templateName string, buffer *bytes.Buffer) {
	var template *template.Template = createTemplate(templateName)
	template.ExecuteTemplate(buffer, "layout", nil)
}

func serveTemplate(templateName string, context *gin.Context) {
	var htmlBuffer bytes.Buffer

	if os.Getenv("GO_ENV") == "dev" {
		executeTemplateDynamic(templateName, &htmlBuffer)
	} else {
		executeTemplateStatic(templateName, &htmlBuffer)
	}

	context.Data(http.StatusOK, "text/html", htmlBuffer.Bytes())
}
