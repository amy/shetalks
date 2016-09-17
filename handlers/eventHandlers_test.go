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

func TestHandlers_AddEvent(t *testing.T) {

	var table = []struct {
		jsonStr       []byte
		expectedEvent shetalks.Event
		createInvoked bool
	}{
		{
			[]byte(`{"name": "test name", "description": "test description", "speakers": [111, 111]}`),
			shetalks.Event{Name: "test name", Description: "test description", Speakers: []int{111, 111}},
			true,
		},
		// add more cases here //
	}

	for _, tb := range table {

		// configure mock
		var es mock.EventService
		es.CreateFn = func(e shetalks.Event) (shetalks.Event, error) {

			fmt.Println("\nHIT HIT HIT\nHIT HIT HIT\nHIT HIT HIT\nHIT HIT HIT\nHIT HIT HIT")

			if !reflect.DeepEqual(e, tb.expectedEvent) {
				t.Errorf("handler created wrong event: got %v want %v", e, tb.expectedEvent)
			}

			return e, nil
		}

		// Invoke AddEvent
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/event", bytes.NewBuffer(tb.jsonStr))
		r.Header.Set("Content-Type", "application/json")

		handlers.AddEvent(&es).ServeHTTP(w, r)

		// Assertions //

		// Check the status code is what we expect.
		/*if status := r.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := `{"alive": true}`
		if r.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		} */

		if es.CreateInvoked != tb.createInvoked {

			if tb.createInvoked == true {
				t.Errorf("expected Create() to be invoked")
			}

			t.Errorf("expected Create() not to be invoked")
		}
	}

}
