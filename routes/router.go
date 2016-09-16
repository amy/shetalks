package routes

import (
	"net/http"

	"github.com/amy/shetalks"
	"github.com/amy/shetalks/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(es shetalks.EventService) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods(http.MethodPost).
		Path("/event").
		Name("AddEvent").
		Handler(handlers.AddEvent(es))
		// should shetalks.AddEvent somehow be injected?
		// & thus should AddEvent be attached to a receiver?

	router.
		Methods(http.MethodGet).
		Path("/event/{id}").
		Name("FindEvent").
		Handler(handlers.FindEvent(es))

	router.
		Methods(http.MethodPut).
		Path("/event/{id}").
		Name("ReplaceEvent").
		Handler(handlers.ReplaceEvent(es))

	router.
		Methods(http.MethodDelete).
		Path("/event/{id}").
		Name("DestroyEvent").
		Handler(handlers.DestroyEvent(es))

	return router
}
