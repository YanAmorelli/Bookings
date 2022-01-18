package handlers

import (
	"net/http"
	"github.com/YanAmorelli/bookings/pkg/config"
	"github.com/YanAmorelli/bookings/pkg/models"
	"github.com/YanAmorelli/bookings/pkg/render"
) 

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	render.RenderTemplate(w, "about.page.go.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}