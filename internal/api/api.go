package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/config"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/store"
	"net/http"
)

type Api struct {
	conf  *config.Configuration
	store store.Store
}

func New(conf *config.Configuration, store store.Store) *Api {
	api := Api{
		conf:  conf,
		store: store,
	}
	return &api
}

func (a *Api) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/health", a.healthHandler).Methods("GET")
	r.HandleFunc("/vibes/", a.addVibe).Methods("POST")
	r.HandleFunc("/vibes/{vibeId}", a.getVibe).Methods("GET")
	r.HandleFunc("/vibes", a.listVibes).Methods("GET").Queries("max", "{max:[0-9]+}")
	r.HandleFunc("/vibes/{vibeId}", a.deleteVibe).Methods("DELETE")

	serverAddress := fmt.Sprintf(":%s", a.conf.ServerPort)
	log.Infof("server started on address %s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}

func (a *Api) writeResponse(w http.ResponseWriter, status int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		a.writeSimpleResponse(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}

func (a *Api) writeSimpleResponse(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))
}

type failedResponse struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func newFailedResponse(status int, detail string) *failedResponse {
	return &failedResponse{
		Status: status,
		Detail: detail,
	}
}

func (a *Api) writeFailedResponse(w http.ResponseWriter, code int, err error) {
	fr := newFailedResponse(code, err.Error())
	a.writeResponse(w, code, fr)
}
