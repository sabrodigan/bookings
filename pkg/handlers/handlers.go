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

// Index is the handler for the default  page
func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	render.RenderTemplate(w, "./templates/index.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["flash"] = "This is dynamic data for the home page"

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Add is the handler for the divide page
func (m *Repository) Add(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "add.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Create is the handler for the divide page
func (m *Repository) Create(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "create.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Delete is the handler for the divide page
func (m *Repository) Delete(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "delete.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Divide is the handler for the divide page
func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "divide.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "home.layout.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Move is the handler for the divide page
func (m *Repository) Move(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "move.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Search is the handler for the divide page
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "search.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
