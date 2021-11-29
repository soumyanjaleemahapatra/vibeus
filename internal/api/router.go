package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/vibe"
	"net/http"
)

func (a *Api) healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("health check")
	a.writeSimpleResponse(w, http.StatusOK)
}

func (a *Api) addVibe(w http.ResponseWriter, r *http.Request) {
	var vb *vibe.Vibe

	err := json.NewDecoder(r.Body).Decode(&vb)
	if err != nil {
		a.writeFailedResponse(w, http.StatusBadRequest, err)
		return
	}

	err = a.store.CreateVibe(r.Context(), vb)
	if err != nil {
		a.writeFailedResponse(w, http.StatusInternalServerError, err)
		return
	}

	a.writeSimpleResponse(w, http.StatusOK)
}

func (a *Api) getVibe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["vibeId"]

	scream, err := a.store.GetVibe(r.Context(), id)
	if err != nil {
		a.writeFailedResponse(w, http.StatusNotFound, err)
		return
	}

	a.writeResponse(w, http.StatusOK, scream)
}

func (a *Api) listVibes(w http.ResponseWriter, r *http.Request) {
	vibes, err := a.store.ListVibes(r.Context())
	if err != nil {
		a.writeFailedResponse(w, http.StatusNotFound, err)
		return
	}
	a.writeResponse(w, http.StatusOK, vibes)
}

func (a *Api) deleteVibe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["vibeId"]

	err := a.store.DeleteVibe(r.Context(), id)
	if err != nil {
		a.writeFailedResponse(w, http.StatusNotFound, err)
		return
	}
	a.writeSimpleResponse(w, http.StatusOK)
}
