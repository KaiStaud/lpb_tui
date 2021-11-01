package handlers

import (
	"log"
	"lpb/data"
	"net/http"
)

type Profiles struct {
	l *log.Logger
}

func NewProfiles(l *log.Logger) *Profiles {
	return &Profiles{l}
}

// implements interface for http handler with type profiles
func (p *Profiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.GetProfiles(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.AddProfile(rw, r)
		return
	}
	// else: catch all
	rw.WriteHeader(http.StatusNotImplemented)
}

// Extract core functionality of getting and marshaling data
func (p *Profiles) GetProfiles(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProfiles()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// Add new Profile to list
func (p *Profiles) AddProfile(rw http.ResponseWriter, r *http.Request) {
	prof := &data.Profile{}
	err := prof.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("PRofile: %#v", prof)
	data.AddProfile(prof)
}
