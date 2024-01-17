package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Insere os templates html na variavel templetes
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// Renderiza uma pagina html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
