package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/scream"
	"net/http"
)

func (a *Api) healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("health check")
	a.writeSimpleResponse(w, http.StatusOK)
}

func (a *Api) addScream(w http.ResponseWriter, r *http.Request) {
	var sc *scream.Scream

	err := json.NewDecoder(r.Body).Decode(&sc)
	if err != nil {
		a.writeFailedResponse(w, http.StatusBadRequest, err)
		return
	}

	err = a.store.CreateScream(r.Context(), sc)
	if err != nil {
		a.writeFailedResponse(w, http.StatusInternalServerError, err)
		return
	}

	a.writeSimpleResponse(w, http.StatusOK)
}

func (a *Api) getScream(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["screamId"]

	scream, err := a.store.GetScream(r.Context(), id)
	if err != nil {
		a.writeFailedResponse(w, http.StatusNotFound, err)
		return
	}

	a.writeResponse(w, http.StatusOK, scream)
}

func (a *Api) listScreams(w http.ResponseWriter, r *http.Request) {
	screams, err := a.store.ListScreams(r.Context())
	if err != nil {
		a.writeFailedResponse(w, http.StatusNotFound, err)
		return
	}
	a.writeResponse(w, http.StatusOK, screams)
}

func (a *Api) deleteScream(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["screamId"]

	err := a.store.DeleteScream(r.Context(), id)
	if err != nil {
		a.writeFailedResponse(w, http.StatusNotFound, err)
		return
	}
	a.writeSimpleResponse(w, http.StatusOK)
}
