package handlers

import (
	"github.com/sabrodigan/bookings/models"
	"github.com/sabrodigan/bookings/pkg/config"
	"github.com/sabrodigan/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	render.PageRender(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Index is the handler for the default  page
func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	render.PageRender(w, "index.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Divide is the handler for the divide page
func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	render.PageRender(w, "divide.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["flash"] = "This is dynamic data for the home page"

	// send data to the template
	render.PageRender(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
