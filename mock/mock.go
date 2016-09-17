package mock

import "github.com/amy/shetalks"

type EventService struct {
	CreateFn      func(e shetalks.Event) (shetalks.Event, error)
	CreateInvoked bool
	ReadFn        func(id int) (shetalks.Event, error)
	ReadInvoked   bool
	UpdateFn      func(id int, name string, description string, speakers []int) (shetalks.Event, error)
	UpdateInvoked bool
	DeleteFn      func(id int) error
	DeleteInvoked bool
}

func (es *EventService) Create(e shetalks.Event) (shetalks.Event, error) {
	es.CreateInvoked = true
	return es.CreateFn(e)
}

func (es *EventService) Read(id int) (shetalks.Event, error) {
	return shetalks.Event{}, nil
}

func (es *EventService) Update(id int, name string, description string, speakers []int) (shetalks.Event, error) {
	return shetalks.Event{}, nil
}

func (es *EventService) Delete(id int) error {
	return nil
}
