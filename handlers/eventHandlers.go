package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/amy/shetalks"
	"github.com/gorilla/mux"
)

func AddEvent(es shetalks.EventService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var e shetalks.Event

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := r.Body.Close(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(body, &e); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity) // unable to write json error to stream w
			}
			return
		}

		e, err = es.Create(e)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity) // unable to write json error to stream w
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(e)
	})

}

func FindEvent(es shetalks.EventService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		e, err := es.Read(id)
		if err != nil {
			// wrong response. Check error to see if it means that there was no
			// entry with the id. Respond with correct response
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(e)
	})

}

func ReplaceEvent(es shetalks.EventService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		var event shetalks.Event

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := r.Body.Close(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(body, &event); err != nil {
		}

		e, err := es.Update(id, event.Name, event.Description, event.Speakers)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(e)
	})

}

func DestroyEvent(es shetalks.EventService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = es.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

}
