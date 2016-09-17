package shetalks

type Event struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Speakers    []int  `json: "speakers"`
}

type EventService interface {
	Create(e Event) (Event, error)
	Read(id int) (Event, error)
	// Change update to accept Event as a param instead
	Update(id int, name string, description string, speakers []int) (Event, error)
	Delete(id int) error
}
