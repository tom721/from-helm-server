package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tom721/from-helm-server/pkg/handler"
)

const (
	port              = 8081
	helmReleasePrefix = "/api/helm/release"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc(helmReleasePrefix, handler.HelmRelease).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
