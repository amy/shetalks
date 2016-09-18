package handlers_test

// testing handlers: https://elithrar.github.io/article/testing-http-handlers-go/

import (
	"bytes"
	"fmt"
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
			[]byte(`{"name":"test name","description":"test description","Speakers":[111,111]}`),
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

		// Assertions //

		// Check the status code is what we expect.

		if status := rr.Code; status != tb.expectedStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tb.expectedStatusCode)
		}

		// Check the response body is what we expect.

		fmt.Printf("Body String: %v", rr.Body.String())

		if rr.Body.String() != `{"name":"test name","description":"test description","Speakers":[111,111]}` {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), `{"name":"test name","description":"test description","Speakers":[111,111]}`)
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
