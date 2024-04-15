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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	//render.TempRendered(w, "home.page.tmpl", &models.TemplateData{})
	//// send data to the template
	stringMap := make(map[string]string)
	render.TempRendered(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Flash:     "This is a flash message",
	})

} // About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["flash"] = "This is dynamic data for the home page"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.TempRendered(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Flash:     "This is a flash message",
	})

}
