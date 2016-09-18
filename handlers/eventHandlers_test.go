package handlers_test

// testing handlers: https://elithrar.github.io/article/testing-http-handlers-go/

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/amy/shetalks"
	"github.com/amy/shetalks/handlers"
	"github.com/amy/shetalks/mock"
)

// @TODO try to break this into subtests

func TestHandlers_AddEvent(t *testing.T) {

	var addEvent = []struct {
		jsonStr            []byte
		expectedEvent      shetalks.Event
		createInvoked      bool
		expectedStatusCode int
	}{
		{
			[]byte(`{"name":"test name","description":"test description","Speakers":[111,111]}` + "\n"),
			shetalks.Event{Name: "test name", Description: "test description", Speakers: []int{111, 111}},
			true,
			http.StatusCreated,
		},
		// add more cases here //
	}

	for _, tb := range addEvent {

		// configure mock
		var es mock.EventService
		es.CreateFn = func(e shetalks.Event) (shetalks.Event, error) {

			if !reflect.DeepEqual(e, tb.expectedEvent) {
				t.Errorf("handler created wrong event: got %v want %v", e, tb.expectedEvent)
			}

			return e, nil
		}

		// Request
		r, _ := http.NewRequest(http.MethodPost, "/event", bytes.NewBuffer(tb.jsonStr))
		r.Header.Set("Content-Type", "application/json")

		// Response Recorder
		rr := httptest.NewRecorder()

		// Invoke AddEvent
		handlers.AddEvent(&es).ServeHTTP(rr, r)

		// Assertions
		if statusCode := rr.Code; statusCode != tb.expectedStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, tb.expectedStatusCode)
		}

		var e shetalks.Event
		json.NewDecoder(rr.Body).Decode(&e)
		if response := e; !reflect.DeepEqual(response, tb.expectedEvent) {
			t.Errorf("handler returned unexpected body: got %q want %q",
				response, tb.expectedEvent)
		}

		if es.CreateInvoked != tb.createInvoked {
			if tb.createInvoked {
				t.Errorf("expected Create() to be invoked")
			} else {
				t.Errorf("expected Create() not to be invoked")
			}
		}
	}

}
