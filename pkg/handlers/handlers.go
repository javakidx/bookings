package handlers

import (
	"net/http"

	"github.com/javakidx/bookings/pkg/config"
	"github.com/javakidx/bookings/pkg/modles"
	"github.com/javakidx/bookings/pkg/render"
)

// Repo the repository used by handlers
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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_id", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &modles.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again. Hahaha"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_id")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &modles.TemplateData{
		StringMap: stringMap,
	})
}
