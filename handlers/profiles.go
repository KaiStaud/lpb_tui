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
