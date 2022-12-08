package handlers

import (
	"net/http"

	"github.com/uwezukwechibuzor/bookings/pkg/config"
	"github.com/uwezukwechibuzor/bookings/pkg/models"
	"github.com/uwezukwechibuzor/bookings/pkg/render"
)

// Repo  is the repositry used by the handlers
var Repo *Repositry

//Repositry is the repositry type
type Repositry struct {
	App *config.AppConfig
}

// NewRepo creates a new repositry
func NewRepo(a *config.AppConfig) *Repositry {
	return &Repositry{
		App: a,
	}
}

// NewHandlers sets the repository of the handlers
func NewHandlers(r *Repositry) {
	Repo = r
}

func (m *Repositry) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

func (m *Repositry) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	StringMap := make(map[string]string)
	StringMap["test"] = "Hello Daniel"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: StringMap,
	})
}
