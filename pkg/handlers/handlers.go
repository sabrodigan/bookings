package handlers

import (
	"net/http"
	"website/models"
	"website/pkg/config"
	"website/pkg/render"
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
	render.TempRendered(w, "home.page.tmpl", &models.TemplateData{})

} // About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again, this is dynamic content, which is sent to the template using the EXECUTE command and including the data in the TemplateData struct.  Once the template has the data, we can show a search page that allows for freetex search of the boxes database."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.TempRendered(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Flash:     "This is a flash message",
	})

}
