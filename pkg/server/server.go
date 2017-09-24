// Package server provides functionality for handling data input and queries
package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/onuryilmaz/boilerplate-go/pkg/commons"
)

// REST provides functionality for HTTP REST API Server
type REST struct {
	router *httprouter.Router
	server *http.Server
	port   string
}

// NewREST creates a REST API server instance with the provided options
func NewREST(options commons.Options) *REST {
	rest := &REST{}
	rest.port = options.ServerPort
	rest.router = httprouter.New()
	return rest
}

// Start starts REST API server and connects handlers to the router on port
func (r *REST) Start() {

	logrus.Info("Starting REST server...")
	logrus.Infof("REST server connecting to port %v", r.port)

	r.router.GET("/ping", r.pingHandler)

	r.server = &http.Server{Addr: ":" + r.port, Handler: r.router}
	go r.server.ListenAndServe()
}

// Stop stops REST API gracefully
func (r *REST) Stop() {
	logrus.Warn("Stopping REST server..")
	r.server.Shutdown(context.TODO())
}

func (r *REST) pingHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	data := map[string]string{"Status": "OK"}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
